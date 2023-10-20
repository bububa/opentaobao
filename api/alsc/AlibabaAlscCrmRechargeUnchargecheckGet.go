package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeUnchargecheckGet 储值账户退充值校验
// alibaba.alsc.crm.recharge.unchargecheck.get
//
// 储值账户退充值校验接口
func AlibabaAlscCrmRechargeUnchargecheckGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeUnchargecheckGetAPIRequest, resp *alsc.AlibabaAlscCrmRechargeUnchargecheckGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
