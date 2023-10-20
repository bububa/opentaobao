package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCustomerVoucherListAPIRequest 获取顾客优惠券列表 API请求
// alibaba.alsc.crm.customer.voucher.list
//
// 获取顾客优惠券列表
type AlibabaAlscCrmCustomerVoucherListAPIRequest struct {
	model.Params
	// 获取顾客优惠券列表
	_customerVoucherFullOpenReq *CustomerVoucherFullOpenReq
}

// NewAlibabaAlscCrmCustomerVoucherListRequest 初始化AlibabaAlscCrmCustomerVoucherListAPIRequest对象
func NewAlibabaAlscCrmCustomerVoucherListRequest() *AlibabaAlscCrmCustomerVoucherListAPIRequest {
	return &AlibabaAlscCrmCustomerVoucherListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCustomerVoucherListAPIRequest) Reset() {
	r._customerVoucherFullOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.voucher.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomerVoucherFullOpenReq is CustomerVoucherFullOpenReq Setter
// 获取顾客优惠券列表
func (r *AlibabaAlscCrmCustomerVoucherListAPIRequest) SetCustomerVoucherFullOpenReq(_customerVoucherFullOpenReq *CustomerVoucherFullOpenReq) error {
	r._customerVoucherFullOpenReq = _customerVoucherFullOpenReq
	r.Set("customer_voucher_full_open_req", _customerVoucherFullOpenReq)
	return nil
}

// GetCustomerVoucherFullOpenReq CustomerVoucherFullOpenReq Getter
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetCustomerVoucherFullOpenReq() *CustomerVoucherFullOpenReq {
	return r._customerVoucherFullOpenReq
}

var poolAlibabaAlscCrmCustomerVoucherListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCustomerVoucherListRequest()
	},
}

// GetAlibabaAlscCrmCustomerVoucherListRequest 从 sync.Pool 获取 AlibabaAlscCrmCustomerVoucherListAPIRequest
func GetAlibabaAlscCrmCustomerVoucherListAPIRequest() *AlibabaAlscCrmCustomerVoucherListAPIRequest {
	return poolAlibabaAlscCrmCustomerVoucherListAPIRequest.Get().(*AlibabaAlscCrmCustomerVoucherListAPIRequest)
}

// ReleaseAlibabaAlscCrmCustomerVoucherListAPIRequest 将 AlibabaAlscCrmCustomerVoucherListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCustomerVoucherListAPIRequest(v *AlibabaAlscCrmCustomerVoucherListAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCustomerVoucherListAPIRequest.Put(v)
}
