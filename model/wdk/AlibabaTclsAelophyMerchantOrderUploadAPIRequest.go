package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantOrderUploadAPIRequest 商家订单数据上传 API请求
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
type AlibabaTclsAelophyMerchantOrderUploadAPIRequest struct {
	model.Params
	// 商家订单信息
	_orderInfo *MerchantOrderInfo
}

// NewAlibabaTclsAelophyMerchantOrderUploadRequest 初始化AlibabaTclsAelophyMerchantOrderUploadAPIRequest对象
func NewAlibabaTclsAelophyMerchantOrderUploadRequest() *AlibabaTclsAelophyMerchantOrderUploadAPIRequest {
	return &AlibabaTclsAelophyMerchantOrderUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderInfo Setter
// 商家订单信息
func (r *AlibabaTclsAelophyMerchantOrderUploadAPIRequest) SetOrderInfo(_orderInfo *MerchantOrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// Get OrderInfo Getter
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetOrderInfo() *MerchantOrderInfo {
	return r._orderInfo
}
