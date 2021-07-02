package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketCardConsumecard 电子凭证储值卡核销
// taobao.vmarket.eticket.card.consumecard
//
// 线下商户核销时，ISV调用电子凭证的isv接口来对电子凭证储值卡核销对应金额
func TaobaoVmarketEticketCardConsumecard(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketCardConsumecardAPIRequest, session string) (*eticket.TaobaoVmarketEticketCardConsumecardAPIResponse, error) {
	var resp eticket.TaobaoVmarketEticketCardConsumecardAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
