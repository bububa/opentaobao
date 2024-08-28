package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeChargeprecheckGet 储值账户充值前校验
// alibaba.alsc.crm.recharge.chargeprecheck.get
//
// 储值账户充值前校验接口
func AlibabaAlscCrmRechargeChargeprecheckGet(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeChargeprecheckGetAPIRequest, resp *alsc.AlibabaAlscCrmRechargeChargeprecheckGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
