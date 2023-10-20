package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeChargeUpdate 储值充值
// alibaba.alsc.crm.recharge.charge.update
//
// 顾客储值账户充值
func AlibabaAlscCrmRechargeChargeUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeChargeUpdateAPIRequest, resp *alsc.AlibabaAlscCrmRechargeChargeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
