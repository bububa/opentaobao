package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeCarEticketConsume 天猫汽车电子凭证核销
// tmall.aliauto.trade.car.eticket.consume
//
// 为商家提供电子凭证核销接口，支持分账
func TmallAliautoTradeCarEticketConsume(clt *core.SDKClient, req *tmallcar.TmallAliautoTradeCarEticketConsumeAPIRequest, resp *tmallcar.TmallAliautoTradeCarEticketConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
