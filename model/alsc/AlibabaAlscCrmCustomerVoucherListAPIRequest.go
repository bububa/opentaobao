package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCustomerVoucherListAPIRequest
获取顾客优惠券列表 API请求
alibaba.alsc.crm.customer.voucher.list

获取顾客优惠券列表 */
type AlibabaAlscCrmCustomerVoucherListAPIRequest struct {
	model.Params
	// 获取顾客优惠券列表
	_customerVoucherFullOpenReq *CustomerVoucherFullOpenReq
}

// NewAlibabaAlscCrmCustomerVoucherListRequest 初始化AlibabaAlscCrmCustomerVoucherListAPIRequest对象
func NewAlibabaAlscCrmCustomerVoucherListRequest() *AlibabaAlscCrmCustomerVoucherListAPIRequest {
	return &AlibabaAlscCrmCustomerVoucherListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.voucher.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CustomerVoucherFullOpenReq Setter
// 获取顾客优惠券列表
func (r *AlibabaAlscCrmCustomerVoucherListAPIRequest) SetCustomerVoucherFullOpenReq(_customerVoucherFullOpenReq *CustomerVoucherFullOpenReq) error {
	r._customerVoucherFullOpenReq = _customerVoucherFullOpenReq
	r.Set("customer_voucher_full_open_req", _customerVoucherFullOpenReq)
	return nil
}

// Get CustomerVoucherFullOpenReq Getter
func (r AlibabaAlscCrmCustomerVoucherListAPIRequest) GetCustomerVoucherFullOpenReq() *CustomerVoucherFullOpenReq {
	return r._customerVoucherFullOpenReq
}
