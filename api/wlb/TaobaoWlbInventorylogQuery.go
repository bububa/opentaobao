package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbInventorylogQuery 根据商品ID查询所有库存变更记录
// taobao.wlb.inventorylog.query
//
// 通过商品ID等几个条件来分页查询库存变更记录
func TaobaoWlbInventorylogQuery(clt *core.SDKClient, req *wlb.TaobaoWlbInventorylogQueryAPIRequest, resp *wlb.TaobaoWlbInventorylogQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
