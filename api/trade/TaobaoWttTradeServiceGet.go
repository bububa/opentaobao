package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoWttTradeServiceGet 获取网厅号卡垂直标信息
// taobao.wtt.trade.service.get
//
// 查询网厅订单信息
func TaobaoWttTradeServiceGet(clt *core.SDKClient, req *trade.TaobaoWttTradeServiceGetAPIRequest, resp *trade.TaobaoWttTradeServiceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
