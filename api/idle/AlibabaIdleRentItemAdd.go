package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemAdd 租赁商品发布
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
func AlibabaIdleRentItemAdd(clt *core.SDKClient, req *idle.AlibabaIdleRentItemAddAPIRequest, resp *idle.AlibabaIdleRentItemAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
