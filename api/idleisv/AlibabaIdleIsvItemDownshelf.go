package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvItemDownshelf 服务商闲鱼商品下架
// alibaba.idle.isv.item.downshelf
//
// 供外部服务商ISV进行闲鱼商品下架操作
func AlibabaIdleIsvItemDownshelf(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemDownshelfAPIRequest, resp *idleisv.AlibabaIdleIsvItemDownshelfAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
