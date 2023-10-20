package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderPageGet 分页查询物流宝订单
// taobao.wlb.order.page.get
//
// 分页查询物流宝订单
func TaobaoWlbOrderPageGet(clt *core.SDKClient, req *wlb.TaobaoWlbOrderPageGetAPIRequest, resp *wlb.TaobaoWlbOrderPageGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
