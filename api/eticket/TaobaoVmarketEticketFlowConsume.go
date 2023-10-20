package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketFlowConsume 无交易类凭证核销
// taobao.vmarket.eticket.flow.consume
//
// 无交易类凭证核销
func TaobaoVmarketEticketFlowConsume(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFlowConsumeAPIRequest, resp *eticket.TaobaoVmarketEticketFlowConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
