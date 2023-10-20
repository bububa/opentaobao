package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantOrderUploadAPIRequest) Reset() {
	r._orderInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderInfo is OrderInfo Setter
// 商家订单信息
func (r *AlibabaTclsAelophyMerchantOrderUploadAPIRequest) SetOrderInfo(_orderInfo *MerchantOrderInfo) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlibabaTclsAelophyMerchantOrderUploadAPIRequest) GetOrderInfo() *MerchantOrderInfo {
	return r._orderInfo
}

var poolAlibabaTclsAelophyMerchantOrderUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantOrderUploadRequest()
	},
}

// GetAlibabaTclsAelophyMerchantOrderUploadRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantOrderUploadAPIRequest
func GetAlibabaTclsAelophyMerchantOrderUploadAPIRequest() *AlibabaTclsAelophyMerchantOrderUploadAPIRequest {
	return poolAlibabaTclsAelophyMerchantOrderUploadAPIRequest.Get().(*AlibabaTclsAelophyMerchantOrderUploadAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantOrderUploadAPIRequest 将 AlibabaTclsAelophyMerchantOrderUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantOrderUploadAPIRequest(v *AlibabaTclsAelophyMerchantOrderUploadAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantOrderUploadAPIRequest.Put(v)
}
