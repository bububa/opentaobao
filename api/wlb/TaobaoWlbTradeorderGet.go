package wlb

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbTradeorderGet 根据交易号获取物流宝订单
// taobao.wlb.tradeorder.get
//
// 根据交易类型和交易id查询物流宝订单详情
func TaobaoWlbTradeorderGet(clt *core.SDKClient, req *wlb.TaobaoWlbTradeorderGetAPIRequest, resp *wlb.TaobaoWlbTradeorderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
