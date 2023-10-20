package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelTradeDeliver 飞猪度假-订单发货接口
// alitrip.travel.trade.deliver
//
// 航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
func AlitripTravelTradeDeliver(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeDeliverAPIRequest, resp *traveltrade.AlitripTravelTradeDeliverAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
