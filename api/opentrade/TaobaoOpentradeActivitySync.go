package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeActivitySync 尖货交易活动信息同步
// taobao.opentrade.activity.sync
//
// 尖货交易活动信息配置，创建或更新活动信息
// 在活动时间开始前，所有用户（包括标记可购买的用户），无法购买商品；
// 在活动时间内，标记可购买的用户可在小程序中跳转下单页，完成购买；
// 在活动结束后，对限购不再限制，平台开放购买，用户可在小程序内、商品详情、购物车下单购买；
func TaobaoOpentradeActivitySync(clt *core.SDKClient, req *opentrade.TaobaoOpentradeActivitySyncAPIRequest, resp *opentrade.TaobaoOpentradeActivitySyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
