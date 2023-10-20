package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) Reset() {
	r._bindingConsignorderGiftRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.gift.binding"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangConsignorderGiftBindingAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangConsignorderGiftBindingRequest()
	},
}

// GetAlibabaDchainAoxiangConsignorderGiftBindingRequest 从 sync.Pool 获取 AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest
func GetAlibabaDchainAoxiangConsignorderGiftBindingAPIRequest() *AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest {
	return poolAlibabaDchainAoxiangConsignorderGiftBindingAPIRequest.Get().(*AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest)
}

// ReleaseAlibabaDchainAoxiangConsignorderGiftBindingAPIRequest 将 AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangConsignorderGiftBindingAPIRequest(v *AlibabaDchainAoxiangConsignorderGiftBindingAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangConsignorderGiftBindingAPIRequest.Put(v)
}
