package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeUndedutUpdate 储值消费退款(逆向)
// alibaba.alsc.crm.recharge.undedut.update
//
// 新增储值消费退款接口
func AlibabaAlscCrmRechargeUndedutUpdate(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeUndedutUpdateAPIRequest, resp *alsc.AlibabaAlscCrmRechargeUndedutUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
