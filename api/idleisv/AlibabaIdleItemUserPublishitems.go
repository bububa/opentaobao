package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleItemUserPublishitems 发布的商品列表
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
func AlibabaIdleItemUserPublishitems(clt *core.SDKClient, req *idleisv.AlibabaIdleItemUserPublishitemsAPIRequest, session string) (*idleisv.AlibabaIdleItemUserPublishitemsAPIResponse, error) {
	var resp idleisv.AlibabaIdleItemUserPublishitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
