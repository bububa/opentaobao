package icbuseller

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSellerVendorServiceProcessAPIRequest
服务商客户关联信息 API请求
alibaba.seller.vendor.service.process

服务商客户关联信息 */
type AlibabaSellerVendorServiceProcessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// NewAlibabaSellerVendorServiceProcessRequest 初始化AlibabaSellerVendorServiceProcessAPIRequest对象
func NewAlibabaSellerVendorServiceProcessRequest() *AlibabaSellerVendorServiceProcessAPIRequest {
	return &AlibabaSellerVendorServiceProcessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.service.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderNum Setter
// order_num
func (r *AlibabaSellerVendorServiceProcessAPIRequest) SetOrderNum(_orderNum string) error {
	r._orderNum = _orderNum
	r.Set("order_num", _orderNum)
	return nil
}

// Get OrderNum Getter
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetOrderNum() string {
	return r._orderNum
}
