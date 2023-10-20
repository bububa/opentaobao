package icbuseller

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSellerVendorServiceProcessAPIRequest 服务商客户关联信息 API请求
// alibaba.seller.vendor.service.process
//
// 服务商客户关联信息
type AlibabaSellerVendorServiceProcessAPIRequest struct {
	model.Params
	// order_num
	_orderNum string
}

// NewAlibabaSellerVendorServiceProcessRequest 初始化AlibabaSellerVendorServiceProcessAPIRequest对象
func NewAlibabaSellerVendorServiceProcessRequest() *AlibabaSellerVendorServiceProcessAPIRequest {
	return &AlibabaSellerVendorServiceProcessAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaSellerVendorServiceProcessAPIRequest) Reset() {
	r._orderNum = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetApiMethodName() string {
	return "alibaba.seller.vendor.service.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNum is OrderNum Setter
// order_num
func (r *AlibabaSellerVendorServiceProcessAPIRequest) SetOrderNum(_orderNum string) error {
	r._orderNum = _orderNum
	r.Set("order_num", _orderNum)
	return nil
}

// GetOrderNum OrderNum Getter
func (r AlibabaSellerVendorServiceProcessAPIRequest) GetOrderNum() string {
	return r._orderNum
}

var poolAlibabaSellerVendorServiceProcessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaSellerVendorServiceProcessRequest()
	},
}

// GetAlibabaSellerVendorServiceProcessRequest 从 sync.Pool 获取 AlibabaSellerVendorServiceProcessAPIRequest
func GetAlibabaSellerVendorServiceProcessAPIRequest() *AlibabaSellerVendorServiceProcessAPIRequest {
	return poolAlibabaSellerVendorServiceProcessAPIRequest.Get().(*AlibabaSellerVendorServiceProcessAPIRequest)
}

// ReleaseAlibabaSellerVendorServiceProcessAPIRequest 将 AlibabaSellerVendorServiceProcessAPIRequest 放入 sync.Pool
func ReleaseAlibabaSellerVendorServiceProcessAPIRequest(v *AlibabaSellerVendorServiceProcessAPIRequest) {
	v.Reset()
	poolAlibabaSellerVendorServiceProcessAPIRequest.Put(v)
}
