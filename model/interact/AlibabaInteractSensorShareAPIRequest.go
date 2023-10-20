package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorshareAPIRequest 分享 API请求
// alibaba.interact.sensor.share
//
// 客户端分享
type AlibabainteractsensorshareAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorshareRequest 初始化AlibabainteractsensorshareAPIRequest对象
func NewAlibabainteractsensorshareRequest() *AlibabainteractsensorshareAPIRequest {
	return &AlibabainteractsensorshareAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorshareAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.share"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorshareAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorshareAPIRequest) GetRawParams() model.Params {
	return r.Params
}
