package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorvibrateAPIRequest 客户端震动 API请求
// alibaba.interact.sensor.vibrate
//
// 客户端震动
type AlibabainteractsensorvibrateAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorvibrateRequest 初始化AlibabainteractsensorvibrateAPIRequest对象
func NewAlibabainteractsensorvibrateRequest() *AlibabainteractsensorvibrateAPIRequest {
	return &AlibabainteractsensorvibrateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorvibrateAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.vibrate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorvibrateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorvibrateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
