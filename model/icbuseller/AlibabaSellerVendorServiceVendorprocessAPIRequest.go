package icbuseller

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSellerVendorServiceVendorprocessAPIRequest) Reset() {
	r._orderNum = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.service.vendorprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSellerVendorServiceVendorprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaSellerVendorServiceVendorprocessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSellerVendorServiceVendorprocessRequest()
	},
}

// GetAlibabaSellerVendorServiceVendorprocessRequest 从 sync.Pool 获取 AlibabaSellerVendorServiceVendorprocessAPIRequest
func GetAlibabaSellerVendorServiceVendorprocessAPIRequest() *AlibabaSellerVendorServiceVendorprocessAPIRequest {
	return poolAlibabaSellerVendorServiceVendorprocessAPIRequest.Get().(*AlibabaSellerVendorServiceVendorprocessAPIRequest)
}

// ReleaseAlibabaSellerVendorServiceVendorprocessAPIRequest 将 AlibabaSellerVendorServiceVendorprocessAPIRequest 放入 sync.Pool
func ReleaseAlibabaSellerVendorServiceVendorprocessAPIRequest(v *AlibabaSellerVendorServiceVendorprocessAPIRequest) {
	v.Reset()
	poolAlibabaSellerVendorServiceVendorprocessAPIRequest.Put(v)
}
