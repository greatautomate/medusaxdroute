package model

import (
	"errors"
	"fmt"
	"math/rand"
	"medusaxd-api/common"
	"medusaxd-api/setting"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var group2model2channels map[string]map[string][]*Channel
var channelsIDM map[int]*Channel
var channelSyncLock sync.RWMutex

func InitChannelCache() {
	if !common.MemoryCacheEnabled {
		return
	}
	newChannelId2channel := make(map[int]*Channel)
	var channels []*Channel
	DB.Where("status = ?", common.ChannelStatusEnabled).Find(&channels)
	for _, channel := range channels {
		newChannelId2channel[channel.Id] = channel
	}
	var abilities []*Ability
	DB.Find(&abilities)
	groups := make(map[string]bool)
	for _, ability := range abilities {
		groups[ability.Group] = true
	}
	newGroup2model2channels := make(map[string]map[string][]*Channel)
	newChannelsIDM := make(map[int]*Channel)
	for group := range groups {
		newGroup2model2channels[group] = make(map[string][]*Channel)
	}
	for _, channel := range channels {
		newChannelsIDM[channel.Id] = channel
		groups := strings.Split(channel.Group, ",")
		for _, group := range groups {
			models := strings.Split(channel.Models, ",")
			for _, model := range models {
				if _, ok := newGroup2model2channels[group][model]; !ok {
					newGroup2model2channels[group][model] = make([]*Channel, 0)
				}
				newGroup2model2channels[group][model] = append(newGroup2model2channels[group][model], channel)
			}
		}
	}

	// sort by priority
	for group, model2channels := range newGroup2model2channels {
		for model, channels := range model2channels {
			sort.Slice(channels, func(i, j int) bool {
				return channels[i].GetPriority() > channels[j].GetPriority()
			})
			newGroup2model2channels[group][model] = channels
		}
	}

	channelSyncLock.Lock()
	group2model2channels = newGroup2model2channels
	channelsIDM = newChannelsIDM
	channelSyncLock.Unlock()
	common.SysLog("channels synced from database")
}

func SyncChannelCache(frequency int) {
	for {
		time.Sleep(time.Duration(frequency) * time.Second)
		common.SysLog("syncing channels from database")
		InitChannelCache()
	}
}

func CacheGetRandomSatisfiedChannel(c *gin.Context, group string, model string, retry int) (*Channel, string, error) {
	var channel *Channel
	var err error
	selectGroup := group
	if group == "auto" {
		if len(setting.AutoGroups) == 0 {
			return nil, selectGroup, errors.New("auto groups is not enabled")
		}
		for _, autoGroup := range setting.AutoGroups {
			if common.DebugEnabled {
				println("autoGroup:", autoGroup)
			}
			channel, _ = getRandomSatisfiedChannel(autoGroup, model, retry)
			if channel == nil {
				continue
			} else {
				c.Set("auto_group", autoGroup)
				selectGroup = autoGroup
				if common.DebugEnabled {
					println("selectGroup:", selectGroup)
				}
				break
			}
		}
	} else {
		channel, err = getRandomSatisfiedChannel(group, model, retry)
		if err != nil {
			return nil, group, err
		}
	}
	if channel == nil {
		return nil, group, errors.New("channel not found")
	}
	return channel, selectGroup, nil
}

func getRandomSatisfiedChannel(group string, model string, retry int) (*Channel, error) {
	if strings.HasPrefix(model, "gpt-4-gizmo") {
		model = "gpt-4-gizmo-*"
	}
	if strings.HasPrefix(model, "gpt-4o-gizmo") {
		model = "gpt-4o-gizmo-*"
	}

	// if memory cache is disabled, get channel directly from database
	if !common.MemoryCacheEnabled {
		return GetRandomSatisfiedChannel(group, model, retry)
	}

	channelSyncLock.RLock()
	channels := group2model2channels[group][model]
	channelSyncLock.RUnlock()

	if len(channels) == 0 {
		return nil, errors.New("channel not found")
	}

	uniquePriorities := make(map[int]bool)
	for _, channel := range channels {
		uniquePriorities[int(channel.GetPriority())] = true
	}
	var sortedUniquePriorities []int
	for priority := range uniquePriorities {
		sortedUniquePriorities = append(sortedUniquePriorities, priority)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedUniquePriorities)))

	if retry >= len(uniquePriorities) {
		retry = len(uniquePriorities) - 1
	}
	targetPriority := int64(sortedUniquePriorities[retry])

	// get the priority for the given retry number
	var targetChannels []*Channel
	for _, channel := range channels {
		if channel.GetPriority() == targetPriority {
			targetChannels = append(targetChannels, channel)
		}
	}

	// 平滑系数
	smoothingFactor := 10
	// Calculate the total weight of all channels up to endIdx
	totalWeight := 0
	for _, channel := range targetChannels {
		totalWeight += channel.GetWeight() + smoothingFactor
	}
	// Generate a random value in the range [0, totalWeight)
	randomWeight := rand.Intn(totalWeight)

	// Find a channel based on its weight
	for _, channel := range targetChannels {
		randomWeight -= channel.GetWeight() + smoothingFactor
		if randomWeight < 0 {
			return channel, nil
		}
	}
	// return null if no channel is not found
	return nil, errors.New("channel not found")
}

func CacheGetChannel(id int) (*Channel, error) {
	if !common.MemoryCacheEnabled {
		return GetChannelById(id, true)
	}
	channelSyncLock.RLock()
	defer channelSyncLock.RUnlock()

	c, ok := channelsIDM[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("当前渠道# %d，已不存在", id))
	}
	return c, nil
}

func CacheUpdateChannelStatus(id int, status int) {
	if !common.MemoryCacheEnabled {
		return
	}
	channelSyncLock.Lock()
	defer channelSyncLock.Unlock()
	if channel, ok := channelsIDM[id]; ok {
		channel.Status = status
	}
}
