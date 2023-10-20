package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorblowAPIRequest 吹气 API请求
// alibaba.interact.sensor.blow
//
// 客户端吹气
type AlibabainteractsensorblowAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorblowRequest 初始化AlibabainteractsensorblowAPIRequest对象
func NewAlibabainteractsensorblowRequest() *AlibabainteractsensorblowAPIRequest {
	return &AlibabainteractsensorblowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorblowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.blow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorblowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorblowAPIRequest) GetRawParams() model.Params {
	return r.Params
}
