package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensoropenwindowAPIRequest 客户端打开新页面 API请求
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
type AlibabainteractsensoropenwindowAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensoropenwindowRequest 初始化AlibabainteractsensoropenwindowAPIRequest对象
func NewAlibabainteractsensoropenwindowRequest() *AlibabainteractsensoropenwindowAPIRequest {
	return &AlibabainteractsensoropenwindowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensoropenwindowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.openwindow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensoropenwindowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensoropenwindowAPIRequest) GetRawParams() model.Params {
	return r.Params
}
