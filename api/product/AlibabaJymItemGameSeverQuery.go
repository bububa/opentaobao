package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemGameSeverQuery 查询商品发布客户端下可用服务器列表
// alibaba.jym.item.game.sever.query
//
// 查询商品发布客户端下可用服务器列表
func AlibabaJymItemGameSeverQuery(clt *core.SDKClient, req *product.AlibabaJymItemGameSeverQueryAPIRequest, resp *product.AlibabaJymItemGameSeverQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
