package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantorderuploadAPIRequest 商家订单数据上传 API请求
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
type AlibabatclsaelophymerchantorderuploadAPIRequest struct {
	model.Params
	// 商家订单信息
	_orderInfo *MerchantOrderInfo
}

// NewAlibabatclsaelophymerchantorderuploadRequest 初始化AlibabatclsaelophymerchantorderuploadAPIRequest对象
func NewAlibabatclsaelophymerchantorderuploadRequest() *AlibabatclsaelophymerchantorderuploadAPIRequest {
	return &AlibabatclsaelophymerchantorderuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantorderuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantorderuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantorderuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 商家订单信息
func (r *AlibabatclsaelophymerchantorderuploadAPIRequest) SetOrderInfo(_orderInfo *MerchantOrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlibabatclsaelophymerchantorderuploadAPIRequest) GetOrderInfo() *MerchantOrderInfo {
	return r._orderInfo
}
