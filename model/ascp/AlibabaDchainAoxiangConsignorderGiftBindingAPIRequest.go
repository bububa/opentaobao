package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest 赠品绑赠计算占用 API请求
// alibaba.dchain.aoxiang.consignorder.gift.binding
//
// 赠品绑赠计算占用
type AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest struct {
	model.Params
	// 赠品绑赠计算占用入参
	_bindingConsignorderGiftRequest *BindingConsignOrderGiftRequest
}

// NewAlibabaDchainAoxiangConsignorderGiftBindingRequest 初始化AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest对象
func NewAlibabaDchainAoxiangConsignorderGiftBindingRequest() *AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest {
	return &AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.gift.binding"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBindingConsignorderGiftRequest is BindingConsignorderGiftRequest Setter
// 赠品绑赠计算占用入参
func (r *AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) SetBindingConsignorderGiftRequest(_bindingConsignorderGiftRequest *BindingConsignOrderGiftRequest) error {
	r._bindingConsignorderGiftRequest = _bindingConsignorderGiftRequest
	r.Set("binding_consignorder_gift_request", _bindingConsignorderGiftRequest)
	return nil
}

// GetBindingConsignorderGiftRequest BindingConsignorderGiftRequest Getter
func (r AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) GetBindingConsignorderGiftRequest() *BindingConsignOrderGiftRequest {
	return r._bindingConsignorderGiftRequest
}
