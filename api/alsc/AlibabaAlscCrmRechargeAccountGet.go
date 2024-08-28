package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeAccountGet 查询储值账户信息
// alibaba.alsc.crm.recharge.account.get
//
// 查询储值账户信息接口
func AlibabaAlscCrmRechargeAccountGet(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeAccountGetAPIRequest, resp *alsc.AlibabaAlscCrmRechargeAccountGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
