package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensoraudioAPIRequest 声音 API请求
// alibaba.interact.sensor.audio
//
// 客户端声音
type AlibabainteractsensoraudioAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensoraudioRequest 初始化AlibabainteractsensoraudioAPIRequest对象
func NewAlibabainteractsensoraudioRequest() *AlibabainteractsensoraudioAPIRequest {
	return &AlibabainteractsensoraudioAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensoraudioAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.audio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensoraudioAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensoraudioAPIRequest) GetRawParams() model.Params {
	return r.Params
}
