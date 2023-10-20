package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderCreate 创建物流宝订单
// taobao.wlb.order.create
//
// 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
func TaobaoWlbOrderCreate(clt *core.SDKClient, req *wlb.TaobaoWlbOrderCreateAPIRequest, resp *wlb.TaobaoWlbOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
