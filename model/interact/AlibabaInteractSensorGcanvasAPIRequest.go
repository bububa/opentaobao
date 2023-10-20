package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorgcanvasAPIRequest gcanvas API请求
// alibaba.interact.sensor.gcanvas
//
// gcanvas 功能
type AlibabainteractsensorgcanvasAPIRequest struct {
	model.Params
}

// NewAlibabainteractsensorgcanvasRequest 初始化AlibabainteractsensorgcanvasAPIRequest对象
func NewAlibabainteractsensorgcanvasRequest() *AlibabainteractsensorgcanvasAPIRequest {
	return &AlibabainteractsensorgcanvasAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractsensorgcanvasAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.sensor.gcanvas"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractsensorgcanvasAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractsensorgcanvasAPIRequest) GetRawParams() model.Params {
	return r.Params
}
