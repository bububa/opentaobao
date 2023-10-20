package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbWlborderGet 根据物流宝订单编号查询物流宝订单概要信息
// taobao.wlb.wlborder.get
//
// 根据物流宝订单编号查询物流宝订单概要信息
func TaobaoWlbWlborderGet(clt *core.SDKClient, req *wlb.TaobaoWlbWlborderGetAPIRequest, resp *wlb.TaobaoWlbWlborderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
