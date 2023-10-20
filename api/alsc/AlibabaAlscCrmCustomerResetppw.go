package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerResetppw 重置支付密码
// alibaba.alsc.crm.customer.resetppw
//
// 重置支付密码
func AlibabaAlscCrmCustomerResetppw(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerResetppwAPIRequest, session string) (*alsc.AlibabaAlscCrmCustomerResetppwAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCustomerResetppwAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
