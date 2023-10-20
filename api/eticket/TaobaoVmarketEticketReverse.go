package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketReverse 电子凭证冲正接口
// taobao.vmarket.eticket.reverse
//
// 电子凭证平台冲正接口
func TaobaoVmarketEticketReverse(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketReverseAPIRequest, resp *eticket.TaobaoVmarketEticketReverseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
