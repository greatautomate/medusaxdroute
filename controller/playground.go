package controller

import (
	"errors"
	"fmt"
	"net/http"
	"medusaxd-api/common"
	"medusaxd-api/constant"
	"medusaxd-api/dto"
	"medusaxd-api/middleware"
	"medusaxd-api/model"
	"medusaxd-api/service"
	"medusaxd-api/setting"
	"time"

	"github.com/gin-gonic/gin"
)

func Playground(c *gin.Context) {
	var openaiErr *dto.OpenAIErrorWithStatusCode

	defer func() {
		if openaiErr != nil {
			c.JSON(openaiErr.StatusCode, gin.H{
				"error": openaiErr.Error,
			})
		}
	}()

	useAccessToken := c.GetBool("use_access_token")
	if useAccessToken {
		openaiErr = service.OpenAIErrorWrapperLocal(errors.New("暂不支持使用 access token"), "access_token_not_supported", http.StatusBadRequest)
		return
	}

	playgroundRequest := &dto.PlayGroundRequest{}
	err := common.UnmarshalBodyReusable(c, playgroundRequest)
	if err != nil {
		openaiErr = service.OpenAIErrorWrapperLocal(err, "unmarshal_request_failed", http.StatusBadRequest)
		return
	}

	if playgroundRequest.Model == "" {
		openaiErr = service.OpenAIErrorWrapperLocal(errors.New("请选择模型"), "model_required", http.StatusBadRequest)
		return
	}
	c.Set("original_model", playgroundRequest.Model)
	group := playgroundRequest.Group
	userGroup := c.GetString("group")

	if group == "" {
		group = userGroup
	} else {
		if !setting.GroupInUserUsableGroups(group) && group != userGroup {
			openaiErr = service.OpenAIErrorWrapperLocal(errors.New("无权访问该分组"), "group_not_allowed", http.StatusForbidden)
			return
		}
		c.Set("group", group)
	}
	c.Set("token_name", "playground-"+group)
	channel, finalGroup, err := model.CacheGetRandomSatisfiedChannel(c, group, playgroundRequest.Model, 0)
	if err != nil {
		message := fmt.Sprintf("当前分组 %s 下对于模型 %s 无可用渠道", finalGroup, playgroundRequest.Model)
		openaiErr = service.OpenAIErrorWrapperLocal(errors.New(message), "get_playground_channel_failed", http.StatusInternalServerError)
		return
	}
	middleware.SetupContextForSelectedChannel(c, channel, playgroundRequest.Model)
	common.SetContextKey(c, constant.ContextKeyRequestStartTime, time.Now())

	// Write user context to ensure acceptUnsetRatio is available
	userId := c.GetInt("id")
	userCache, err := model.GetUserCache(userId)
	if err != nil {
		openaiErr = service.OpenAIErrorWrapperLocal(err, "get_user_cache_failed", http.StatusInternalServerError)
		return
	}
	userCache.WriteContext(c)
	Relay(c)
}
