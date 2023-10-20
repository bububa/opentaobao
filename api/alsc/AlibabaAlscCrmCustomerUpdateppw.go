package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerUpdateppw 修改支付密码
// alibaba.alsc.crm.customer.updateppw
//
// 修改支付密码
func AlibabaAlscCrmCustomerUpdateppw(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerUpdateppwAPIRequest, resp *alsc.AlibabaAlscCrmCustomerUpdateppwAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
