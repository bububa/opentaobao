package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbInventoryDetailGet 查询库存详情
// taobao.wlb.inventory.detail.get
//
// 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
func TaobaoWlbInventoryDetailGet(clt *core.SDKClient, req *wlb.TaobaoWlbInventoryDetailGetAPIRequest, resp *wlb.TaobaoWlbInventoryDetailGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
