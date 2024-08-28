package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeAccountFlowdetailGet 储值流水详细
// alibaba.alsc.crm.recharge.account.flowdetail.get
//
// 查询储值流水详细接口
func AlibabaAlscCrmRechargeAccountFlowdetailGet(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeAccountFlowdetailGetAPIRequest, resp *alsc.AlibabaAlscCrmRechargeAccountFlowdetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
