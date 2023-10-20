package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketSend 商家电子凭证发码成功回调接口
// taobao.vmarket.eticket.send
//
// 外部商家成功发码回调接口
func TaobaoVmarketEticketSend(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketSendAPIRequest, resp *eticket.TaobaoVmarketEticketSendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
