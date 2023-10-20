package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorshakeAPIRequest 摇一摇 API请求
// alibaba.interact.sensor.shake
//
// 摇一摇
type AlibabainteractsensorshakeAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorshakeRequest 初始化AlibabainteractsensorshakeAPIRequest对象
func NewAlibabainteractsensorshakeRequest() *AlibabainteractsensorshakeAPIRequest {
	return &AlibabainteractsensorshakeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorshakeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.shake"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorshakeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorshakeAPIRequest) GetRawParams() model.Params {
	return r.Params
}
