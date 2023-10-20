package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangconsignordergiftbindingAPIRequest 赠品绑赠计算占用 API请求
// alibaba.dchain.aoxiang.consignorder.gift.binding
//
// 赠品绑赠计算占用
type AlibabadchainaoxiangconsignordergiftbindingAPIRequest struct {
	model.Params
	// 赠品绑赠计算占用入参
	_bindingConsignorderGiftRequest *BindingConsignOrderGiftRequest
}

// NewAlibabadchainaoxiangconsignordergiftbindingRequest 初始化AlibabadchainaoxiangconsignordergiftbindingAPIRequest对象
func NewAlibabadchainaoxiangconsignordergiftbindingRequest() *AlibabadchainaoxiangconsignordergiftbindingAPIRequest {
	return &AlibabadchainaoxiangconsignordergiftbindingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangconsignordergiftbindingAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.gift.binding"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangconsignordergiftbindingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangconsignordergiftbindingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindingConsignorderGiftRequest is BindingConsignorderGiftRequest Setter
// 赠品绑赠计算占用入参
func (r *AlibabadchainaoxiangconsignordergiftbindingAPIRequest) SetBindingConsignorderGiftRequest(_bindingConsignorderGiftRequest *BindingConsignOrderGiftRequest) error {
	r._bindingConsignorderGiftRequest = _bindingConsignorderGiftRequest
	r.Set("binding_consignorder_gift_request", _bindingConsignorderGiftRequest)
	return nil
}

// GetBindingConsignorderGiftRequest BindingConsignorderGiftRequest Getter
func (r AlibabadchainaoxiangconsignordergiftbindingAPIRequest) GetBindingConsignorderGiftRequest() *BindingConsignOrderGiftRequest {
	return r._bindingConsignorderGiftRequest
}
