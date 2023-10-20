package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketAuthConsume 核销放行的核销接口
// taobao.vmarket.eticket.auth.consume
//
// 针对O2O电子凭证核销放行业务，为满足码商能够核销淘宝码而开放的核销接口
func TaobaoVmarketEticketAuthConsume(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketAuthConsumeAPIRequest, resp *eticket.TaobaoVmarketEticketAuthConsumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
