package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// AlitripTravelTradeQuery 飞猪度假-订单详情查询接口
// alitrip.travel.trade.query
//
// 飞猪度假订单详情查询接口
func AlitripTravelTradeQuery(clt *core.SDKClient, req *traveltrade.AlitripTravelTradeQueryAPIRequest, resp *traveltrade.AlitripTravelTradeQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
