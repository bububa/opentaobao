package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderstatusGet 物流宝订单流转状态查询
// taobao.wlb.orderstatus.get
//
// 根据物流宝订单号查询物流宝订单至目前为止的流转状态列表
func TaobaoWlbOrderstatusGet(clt *core.SDKClient, req *wlb.TaobaoWlbOrderstatusGetAPIRequest, resp *wlb.TaobaoWlbOrderstatusGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
