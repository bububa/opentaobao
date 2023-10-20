package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvItemRechargeEdit 闲鱼商品直充功能编辑
// alibaba.idle.isv.item.recharge.edit
//
// 闲鱼商品直充功能编辑
func AlibabaIdleIsvItemRechargeEdit(clt *core.SDKClient, req *idle.AlibabaIdleIsvItemRechargeEditAPIRequest, session string) (*idle.AlibabaIdleIsvItemRechargeEditAPIResponse, error) {
	var resp idle.AlibabaIdleIsvItemRechargeEditAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
