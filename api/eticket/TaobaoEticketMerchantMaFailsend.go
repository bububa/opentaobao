package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaFailsend 码商发码失败回调接口
// taobao.eticket.merchant.ma.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
func TaobaoEticketMerchantMaFailsend(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaFailsendAPIRequest, resp *eticket.TaobaoEticketMerchantMaFailsendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
