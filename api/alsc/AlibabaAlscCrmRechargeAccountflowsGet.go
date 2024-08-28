package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeAccountflowsGet 分页查询储值流水
// alibaba.alsc.crm.recharge.accountflows.get
//
// 增加分页查询储值流水接口
func AlibabaAlscCrmRechargeAccountflowsGet(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeAccountflowsGetAPIRequest, resp *alsc.AlibabaAlscCrmRechargeAccountflowsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
