package traveltrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/traveltrade"
)

// TaobaoAlitripTravelTradeMemoUpdate 修改一笔交易备注
// taobao.alitrip.travel.trade.memo.update
//
// 更新一笔交易备注
func TaobaoAlitripTravelTradeMemoUpdate(clt *core.SDKClient, req *traveltrade.TaobaoAlitripTravelTradeMemoUpdateAPIRequest, resp *traveltrade.TaobaoAlitripTravelTradeMemoUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
