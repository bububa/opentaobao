package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderitemPageGet 分页查询物流宝订单商品详情
// taobao.wlb.orderitem.page.get
//
// 分页查询物流宝订单商品详情
func TaobaoWlbOrderitemPageGet(clt *core.SDKClient, req *wlb.TaobaoWlbOrderitemPageGetAPIRequest, resp *wlb.TaobaoWlbOrderitemPageGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
