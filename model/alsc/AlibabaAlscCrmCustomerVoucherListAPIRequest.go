package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomervoucherlistAPIRequest 获取顾客优惠券列表 API请求
// alibaba.alsc.crm.customer.voucher.list
//
// 获取顾客优惠券列表
type AlibabaalsccrmcustomervoucherlistAPIRequest struct {
	model.Params
	// 获取顾客优惠券列表
	_customerVoucherFullOpenReq *CustomerVoucherFullOpenReq
}

// NewAlibabaalsccrmcustomervoucherlistRequest 初始化AlibabaalsccrmcustomervoucherlistAPIRequest对象
func NewAlibabaalsccrmcustomervoucherlistRequest() *AlibabaalsccrmcustomervoucherlistAPIRequest {
	return &AlibabaalsccrmcustomervoucherlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcustomervoucherlistAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.voucher.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcustomervoucherlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcustomervoucherlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomerVoucherFullOpenReq is CustomerVoucherFullOpenReq Setter
// 获取顾客优惠券列表
func (r *AlibabaalsccrmcustomervoucherlistAPIRequest) SetCustomerVoucherFullOpenReq(_customerVoucherFullOpenReq *CustomerVoucherFullOpenReq) error {
	r._customerVoucherFullOpenReq = _customerVoucherFullOpenReq
	r.Set("customer_voucher_full_open_req", _customerVoucherFullOpenReq)
	return nil
}

// GetCustomerVoucherFullOpenReq CustomerVoucherFullOpenReq Getter
func (r AlibabaalsccrmcustomervoucherlistAPIRequest) GetCustomerVoucherFullOpenReq() *CustomerVoucherFullOpenReq {
	return r._customerVoucherFullOpenReq
}
