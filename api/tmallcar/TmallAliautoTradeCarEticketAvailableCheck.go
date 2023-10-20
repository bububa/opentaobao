package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeCarEticketAvailableCheck 天猫汽车电子凭证核销前校验
// tmall.aliauto.trade.car.eticket.available.check
//
// 天猫汽车核销码可用性校验
func TmallAliautoTradeCarEticketAvailableCheck(clt *core.SDKClient, req *tmallcar.TmallAliautoTradeCarEticketAvailableCheckAPIRequest, resp *tmallcar.TmallAliautoTradeCarEticketAvailableCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
