package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest
天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 API请求
alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update

天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 */
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest struct {
	model.Params
	// 一次请求的唯一标识符
	_seqId string
	// 新的 key
	_secret string
}

// New
