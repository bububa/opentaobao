package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsOrderGet 物流订单获取
// taobao.wlb.imports.order.get
//
// 一般进口物流订单获取
func TaobaoWlbImportsOrderGet(clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsOrderGetAPIRequest, resp *wlbimports.TaobaoWlbImportsOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
