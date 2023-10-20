package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenCustomerGet 查询会员资产
// alibaba.alsc.crm.open.customer.get
//
// 查询会员身份，资产等
func AlibabaAlscCrmOpenCustomerGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenCustomerGetAPIRequest, resp *alsc.AlibabaAlscCrmOpenCustomerGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
