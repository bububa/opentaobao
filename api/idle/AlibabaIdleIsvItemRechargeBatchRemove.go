package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvItemRechargeBatchRemove 闲鱼商品直充功能移除
// alibaba.idle.isv.item.recharge.batch.remove
//
// 闲鱼商品直充功能移除
func AlibabaIdleIsvItemRechargeBatchRemove(clt *core.SDKClient, req *idle.AlibabaIdleIsvItemRechargeBatchRemoveAPIRequest, resp *idle.AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
