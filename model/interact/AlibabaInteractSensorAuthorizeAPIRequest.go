package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorauthorizeAPIRequest 客户端授权页 API请求
// alibaba.interact.sensor.authorize
//
// 客户端授权页
type AlibabainteractsensorauthorizeAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorauthorizeRequest 初始化AlibabainteractsensorauthorizeAPIRequest对象
func NewAlibabainteractsensorauthorizeRequest() *AlibabainteractsensorauthorizeAPIRequest {
	return &AlibabainteractsensorauthorizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorauthorizeAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.authorize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorauthorizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorauthorizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}
