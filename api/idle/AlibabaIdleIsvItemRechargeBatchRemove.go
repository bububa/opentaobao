package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvItemRechargeBatchRemove 闲鱼商品直充功能移除
// alibaba.idle.isv.item.recharge.batch.remove
//
// 闲鱼商品直充功能移除
func AlibabaIdleIsvItemRechargeBatchRemove(clt *core.SDKClient, req *idle.AlibabaIdleIsvItemRechargeBatchRemoveAPIRequest, session string) (*idle.AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse, error) {
	var resp idle.AlibabaIdleIsvItemRechargeBatchRemoveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
