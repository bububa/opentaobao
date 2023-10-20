package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoTradeCarEticketConsume 天猫汽车电子凭证核销
// tmall.aliauto.trade.car.eticket.consume
//
// 为商家提供电子凭证核销接口，支持分账
func TmallAliautoTradeCarEticketConsume(clt *core.SDKClient, req *tmallcar.TmallAliautoTradeCarEticketConsumeAPIRequest, session string) (*tmallcar.TmallAliautoTradeCarEticketConsumeAPIResponse, error) {
	var resp tmallcar.TmallAliautoTradeCarEticketConsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
