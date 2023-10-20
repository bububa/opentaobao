package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorpopwindowAPIRequest popwindow API请求
// alibaba.interact.sensor.popwindow
//
// popwindow
type AlibabainteractsensorpopwindowAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorpopwindowRequest 初始化AlibabainteractsensorpopwindowAPIRequest对象
func NewAlibabainteractsensorpopwindowRequest() *AlibabainteractsensorpopwindowAPIRequest {
	return &AlibabainteractsensorpopwindowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorpopwindowAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.popwindow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorpopwindowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorpopwindowAPIRequest) GetRawParams() model.Params {
	return r.Params
}
