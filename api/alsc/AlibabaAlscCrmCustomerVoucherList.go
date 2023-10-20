package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerVoucherList 获取顾客优惠券列表
// alibaba.alsc.crm.customer.voucher.list
//
// 获取顾客优惠券列表
func AlibabaAlscCrmCustomerVoucherList(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerVoucherListAPIRequest, resp *alsc.AlibabaAlscCrmCustomerVoucherListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
