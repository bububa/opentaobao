package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemQuery 查询租赁商品信息
// alibaba.idle.rent.item.query
//
// 查询租赁商品信息
func AlibabaIdleRentItemQuery(clt *core.SDKClient, req *idle.AlibabaIdleRentItemQueryAPIRequest, resp *idle.AlibabaIdleRentItemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
