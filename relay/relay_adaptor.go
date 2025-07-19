package relay

import (
	"medusaxd-api/constant"
	commonconstant "medusaxd-api/constant"
	"medusaxd-api/relay/channel"
	"medusaxd-api/relay/channel/ali"
	"medusaxd-api/relay/channel/aws"
	"medusaxd-api/relay/channel/baidu"
	"medusaxd-api/relay/channel/baidu_v2"
	"medusaxd-api/relay/channel/claude"
	"medusaxd-api/relay/channel/cloudflare"
	"medusaxd-api/relay/channel/cohere"
	"medusaxd-api/relay/channel/coze"
	"medusaxd-api/relay/channel/deepseek"
	"medusaxd-api/relay/channel/dify"
	"medusaxd-api/relay/channel/gemini"
	"medusaxd-api/relay/channel/jina"
	"medusaxd-api/relay/channel/mistral"
	"medusaxd-api/relay/channel/mokaai"
	"medusaxd-api/relay/channel/ollama"
	"medusaxd-api/relay/channel/openai"
	"medusaxd-api/relay/channel/palm"
	"medusaxd-api/relay/channel/perplexity"
	"medusaxd-api/relay/channel/siliconflow"
	"medusaxd-api/relay/channel/task/jimeng"
	"medusaxd-api/relay/channel/task/kling"
	"medusaxd-api/relay/channel/task/suno"
	"medusaxd-api/relay/channel/tencent"
	"medusaxd-api/relay/channel/vertex"
	"medusaxd-api/relay/channel/volcengine"
	"medusaxd-api/relay/channel/xai"
	"medusaxd-api/relay/channel/xunfei"
	"medusaxd-api/relay/channel/zhipu"
	"medusaxd-api/relay/channel/zhipu_4v"
)

func GetAdaptor(apiType int) channel.Adaptor {
	switch apiType {
	case constant.APITypeAli:
		return &ali.Adaptor{}
	case constant.APITypeAnthropic:
		return &claude.Adaptor{}
	case constant.APITypeBaidu:
		return &baidu.Adaptor{}
	case constant.APITypeGemini:
		return &gemini.Adaptor{}
	case constant.APITypeOpenAI:
		return &openai.Adaptor{}
	case constant.APITypePaLM:
		return &palm.Adaptor{}
	case constant.APITypeTencent:
		return &tencent.Adaptor{}
	case constant.APITypeXunfei:
		return &xunfei.Adaptor{}
	case constant.APITypeZhipu:
		return &zhipu.Adaptor{}
	case constant.APITypeZhipuV4:
		return &zhipu_4v.Adaptor{}
	case constant.APITypeOllama:
		return &ollama.Adaptor{}
	case constant.APITypePerplexity:
		return &perplexity.Adaptor{}
	case constant.APITypeAws:
		return &aws.Adaptor{}
	case constant.APITypeCohere:
		return &cohere.Adaptor{}
	case constant.APITypeDify:
		return &dify.Adaptor{}
	case constant.APITypeJina:
		return &jina.Adaptor{}
	case constant.APITypeCloudflare:
		return &cloudflare.Adaptor{}
	case constant.APITypeSiliconFlow:
		return &siliconflow.Adaptor{}
	case constant.APITypeVertexAi:
		return &vertex.Adaptor{}
	case constant.APITypeMistral:
		return &mistral.Adaptor{}
	case constant.APITypeDeepSeek:
		return &deepseek.Adaptor{}
	case constant.APITypeMokaAI:
		return &mokaai.Adaptor{}
	case constant.APITypeVolcEngine:
		return &volcengine.Adaptor{}
	case constant.APITypeBaiduV2:
		return &baidu_v2.Adaptor{}
	case constant.APITypeOpenRouter:
		return &openai.Adaptor{}
	case constant.APITypeXinference:
		return &openai.Adaptor{}
	case constant.APITypeXai:
		return &xai.Adaptor{}
	case constant.APITypeCoze:
		return &coze.Adaptor{}
	}
	return nil
}

func GetTaskAdaptor(platform commonconstant.TaskPlatform) channel.TaskAdaptor {
	switch platform {
	//case constant.APITypeAIProxyLibrary:
	//	return &aiproxy.Adaptor{}
	case commonconstant.TaskPlatformSuno:
		return &suno.TaskAdaptor{}
	case commonconstant.TaskPlatformKling:
		return &kling.TaskAdaptor{}
	case commonconstant.TaskPlatformJimeng:
		return &jimeng.TaskAdaptor{}
	}
	return nil
}
