package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 API请求
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
type AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest struct {
	model.Params
	// 新的 key
	_secret string
	// 一次请求的唯一标识符
	_seqId string
}

// NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateRequest 初始化AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest对象
func NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateRequest() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest {
	return &AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) Reset() {
	r._secret = ""
	r._seqId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecret is Secret Setter
// 新的 key
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) SetSecret(_secret string) error {
	r._secret = _secret
	r.Set("secret", _secret)
	return nil
}

// GetSecret Secret Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetSecret() string {
	return r._secret
}

// SetSeqId is SeqId Setter
// 一次请求的唯一标识符
func (r *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) SetSeqId(_seqId string) error {
	r._seqId = _seqId
	r.Set("seq_id", _seqId)
	return nil
}

// GetSeqId SeqId Getter
func (r AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) GetSeqId() string {
	return r._seqId
}

var poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateRequest()
	},
}

// GetAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateRequest 从 sync.Pool 获取 AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest
func GetAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest() *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest {
	return poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest.Get().(*AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest)
}

// ReleaseAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest 将 AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest(v *AlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAliyunAicloudIotVisionSaasCtccJiangsuKeySecretUpdateAPIRequest.Put(v)
}
