package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeUnchargeUpdate 充值退款
// alibaba.alsc.crm.recharge.uncharge.update
//
// 充值退款
func AlibabaAlscCrmRechargeUnchargeUpdate(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeUnchargeUpdateAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRechargeUnchargeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
