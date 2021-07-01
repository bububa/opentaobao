package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractMediaAudioAPIRequest
音频相关鉴权接口 API请求
alibaba.interact.media.audio

新音频包的鉴权接口 */
type AlibabaInteractMediaAudioAPIRequest struct {
	model.Params
	// 系统自动生成
	_id string
}

// NewAlibabaInteractMediaAudioRequest 初始化AlibabaInteractMediaAudioAPIRequest对象
func NewAlibabaInteractMediaAudioRequest() *AlibabaInteractMediaAudioAPIRequest {
	return &AlibabaInteractMediaAudioAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractMediaAudioAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.media.audio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractMediaAudioAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 系统自动生成
func (r *AlibabaInteractMediaAudioAPIRequest) SetId(_id string) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r AlibabaInteractMediaAudioAPIRequest) GetId() string {
	return r._id
}
