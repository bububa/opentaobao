package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新 API请求
// alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update
//
// 天猫精灵 IoT 视频 SaaS 服务-江苏电信-appKeySecret 更新
type AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest struct {
	model.Params
	// 新的 key
	_secret string
	// 一次请求的唯一标识符
	_seqId string
}

// NewAlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateRequest 初始化AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest对象
func NewAlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateRequest() *AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest {
	return &AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.aliyun.aicloud.iot.vision.saas.ctcc.jiangsu.key.secret.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSecret is Secret Setter
// 新的 key
func (r *AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) SetSecret(_secret string) error {
	r._secret = _secret
	r.Set("secret", _secret)
	return nil
}

// GetSecret Secret Getter
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) GetSecret() string {
	return r._secret
}

// SetSeqId is SeqId Setter
// 一次请求的唯一标识符
func (r *AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) SetSeqId(_seqId string) error {
	r._seqId = _seqId
	r.Set("seq_id", _seqId)
	return nil
}

// GetSeqId SeqId Getter
func (r AlibabaaliyunaicloudiotvisionsaasctccjiangsukeysecretupdateAPIRequest) GetSeqId() string {
	return r._seqId
}
