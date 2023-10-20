package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorloginAPIRequest 获取登陆页面 API请求
// alibaba.interact.sensor.login
//
// 获取登陆页面
type AlibabainteractsensorloginAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorloginRequest 初始化AlibabainteractsensorloginAPIRequest对象
func NewAlibabainteractsensorloginRequest() *AlibabainteractsensorloginAPIRequest {
	return &AlibabainteractsensorloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorloginAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.login"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}
