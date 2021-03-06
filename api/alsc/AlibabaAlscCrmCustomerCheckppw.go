package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerCheckppw 校验支付密码
// alibaba.alsc.crm.customer.checkppw
//
// 校验支付密码
func AlibabaAlscCrmCustomerCheckppw(clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerCheckppwAPIRequest, session string) (*alsc.AlibabaAlscCrmCustomerCheckppwAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmCustomerCheckppwAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
