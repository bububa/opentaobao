package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// Alibabaidleitemuserpublishitems 发布的商品列表
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
func Alibabaidleitemuserpublishitems(clt *core.SDKClient, req *idleisv.AlibabaidleitemuserpublishitemsAPIRequest, session string) (*idleisv.AlibabaidleitemuserpublishitemsAPIResponse, error) {
	var resp idleisv.AlibabaidleitemuserpublishitemsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
