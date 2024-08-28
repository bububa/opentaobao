package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketFailsend 无法发码回调
// taobao.vmarket.eticket.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
func TaobaoVmarketEticketFailsend(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFailsendAPIRequest, resp *eticket.TaobaoVmarketEticketFailsendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
