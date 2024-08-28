package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvItemRechargeEdit 闲鱼商品直充功能编辑
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
func AlibabaIdleIsvItemRechargeEdit(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleIsvItemRechargeEditAPIRequest, resp *idle.AlibabaIdleIsvItemRechargeEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
