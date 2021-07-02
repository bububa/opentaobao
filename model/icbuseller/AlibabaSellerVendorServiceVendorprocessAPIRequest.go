package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorServiceVendorprocessAPIRequest 服务商客户关联信息 API请求
// alibaba.seller.vendor.service.vendorprocess
//
// 服务商客户关联信息
type AlibabaSellerVendorServiceVendorprocessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// NewAlibabaSellerVendorServiceVendorprocessRequest 初始化AlibabaSellerVendorServiceVendorprocessAPIRequest对象
func NewAlibabaSellerVendorServiceVendorprocessRequest() *AlibabaSellerVendorServiceVendorprocessAPIRequest {
	return &AlibabaSellerVendorServiceVendorprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.service.vendorprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderNum is OrderNum Setter
// order_num
func (r *AlibabaSellerVendorServiceVendorprocessAPIRequest) SetOrderNum(_orderNum string) error {
	r._orderNum = _orderNum
	r.Set("order_num", _orderNum)
	return nil
}

// GetOrderNum OrderNum Getter
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetOrderNum() string {
	return r._orderNum
}
