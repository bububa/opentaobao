package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensormaAPIRequest 码相关API API请求
// alibaba.interact.sensor.ma
//
// 码相关API
type AlibabainteractsensormaAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensormaRequest 初始化AlibabainteractsensormaAPIRequest对象
func NewAlibabainteractsensormaRequest() *AlibabainteractsensormaAPIRequest {
	return &AlibabainteractsensormaAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensormaAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.ma"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensormaAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensormaAPIRequest) GetRawParams() model.Params {
	return r.Params
}
