package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvItemRechargeEdit 闲鱼商品直充功能编辑
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
func AlibabaIdleIsvItemRechargeEdit(clt *core.SDKClient, req *idle.AlibabaIdleIsvItemRechargeEditAPIRequest, resp *idle.AlibabaIdleIsvItemRechargeEditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
