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

// New
