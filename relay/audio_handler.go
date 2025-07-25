package relay

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"medusaxd-api/common"
	"medusaxd-api/dto"
	relaycommon "medusaxd-api/relay/common"
	relayconstant "medusaxd-api/relay/constant"
	"medusaxd-api/relay/helper"
	"medusaxd-api/service"
	"medusaxd-api/setting"
	"strings"
)

func getAndValidAudioRequest(c *gin.Context, info *relaycommon.RelayInfo) (*dto.AudioRequest, error) {
	audioRequest := &dto.AudioRequest{}
	err := common.UnmarshalBodyReusable(c, audioRequest)
	if err != nil {
		return nil, err
	}
	switch info.RelayMode {
	case relayconstant.RelayModeAudioSpeech:
		if audioRequest.Model == "" {
			return nil, errors.New("model is required")
		}
		if setting.ShouldCheckPromptSensitive() {
			words, err := service.CheckSensitiveInput(audioRequest.Input)
			if err != nil {
				common.LogWarn(c, fmt.Sprintf("user sensitive words detected: %s", strings.Join(words, ",")))
				return nil, err
			}
		}
	default:
		err = c.Request.ParseForm()
		if err != nil {
			return nil, err
		}
		formData := c.Request.PostForm
		if audioRequest.Model == "" {
			audioRequest.Model = formData.Get("model")
		}

		if audioRequest.Model == "" {
			return nil, errors.New("model is required")
		}
		audioRequest.ResponseFormat = formData.Get("response_format")
		if audioRequest.ResponseFormat == "" {
			audioRequest.ResponseFormat = "json"
		}
	}
	return audioRequest, nil
}

func AudioHelper(c *gin.Context) (openaiErr *dto.OpenAIErrorWithStatusCode) {
	relayInfo := relaycommon.GenRelayInfoOpenAIAudio(c)
	audioRequest, err := getAndValidAudioRequest(c, relayInfo)

	if err != nil {
		common.LogError(c, fmt.Sprintf("getAndValidAudioRequest failed: %s", err.Error()))
		return service.OpenAIErrorWrapper(err, "invalid_audio_request", http.StatusBadRequest)
	}

	promptTokens := 0
	preConsumedTokens := common.PreConsumedQuota
	if relayInfo.RelayMode == relayconstant.RelayModeAudioSpeech {
		promptTokens = service.CountTTSToken(audioRequest.Input, audioRequest.Model)
		preConsumedTokens = promptTokens
		relayInfo.PromptTokens = promptTokens
	}

	priceData, err := helper.ModelPriceHelper(c, relayInfo, preConsumedTokens, 0)
	if err != nil {
		return service.OpenAIErrorWrapperLocal(err, "model_price_error", http.StatusInternalServerError)
	}

	preConsumedQuota, userQuota, openaiErr := preConsumeQuota(c, priceData.ShouldPreConsumedQuota, relayInfo)
	if openaiErr != nil {
		return openaiErr
	}
	defer func() {
		if openaiErr != nil {
			returnPreConsumedQuota(c, relayInfo, userQuota, preConsumedQuota)
		}
	}()

	err = helper.ModelMappedHelper(c, relayInfo, audioRequest)
	if err != nil {
		return service.OpenAIErrorWrapperLocal(err, "model_mapped_error", http.StatusInternalServerError)
	}

	adaptor := GetAdaptor(relayInfo.ApiType)
	if adaptor == nil {
		return service.OpenAIErrorWrapperLocal(fmt.Errorf("invalid api type: %d", relayInfo.ApiType), "invalid_api_type", http.StatusBadRequest)
	}
	adaptor.Init(relayInfo)

	ioReader, err := adaptor.ConvertAudioRequest(c, relayInfo, *audioRequest)
	if err != nil {
		return service.OpenAIErrorWrapperLocal(err, "convert_request_failed", http.StatusInternalServerError)
	}

	resp, err := adaptor.DoRequest(c, relayInfo, ioReader)
	if err != nil {
		return service.OpenAIErrorWrapper(err, "do_request_failed", http.StatusInternalServerError)
	}
	statusCodeMappingStr := c.GetString("status_code_mapping")

	var httpResp *http.Response
	if resp != nil {
		httpResp = resp.(*http.Response)
		if httpResp.StatusCode != http.StatusOK {
			openaiErr = service.RelayErrorHandler(httpResp, false)
			// reset status code 重置状态码
			service.ResetStatusCode(openaiErr, statusCodeMappingStr)
			return openaiErr
		}
	}

	usage, openaiErr := adaptor.DoResponse(c, httpResp, relayInfo)
	if openaiErr != nil {
		// reset status code 重置状态码
		service.ResetStatusCode(openaiErr, statusCodeMappingStr)
		return openaiErr
	}

	postConsumeQuota(c, relayInfo, usage.(*dto.Usage), preConsumedQuota, userQuota, priceData, "")

	return nil
}
