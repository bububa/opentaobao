package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmRechargeAccountFlowdetailGet
储值流水详细
alibaba.alsc.crm.recharge.account.flowdetail.get

查询储值流水详细接口 */
func AlibabaAlscCrmRechargeAccountFlowdetailGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest, session string) (*alsc.AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
