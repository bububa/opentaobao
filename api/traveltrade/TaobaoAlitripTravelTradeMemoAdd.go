package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// TaobaoAlitripTravelTradeMemoAdd 添加一笔交易备注
// taobao.alitrip.travel.trade.memo.add
//
// 对一笔交易添加备注
func TaobaoAlitripTravelTradeMemoAdd(clt *core.SDKClient, req *traveltrade.TaobaoAlitripTravelTradeMemoAddAPIRequest, resp *traveltrade.TaobaoAlitripTravelTradeMemoAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
