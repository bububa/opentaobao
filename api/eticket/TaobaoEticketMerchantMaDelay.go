package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaDelay 凭证延期
// taobao.eticket.merchant.ma.delay
//
// 订单延期
func TaobaoEticketMerchantMaDelay(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaDelayAPIRequest, resp *eticket.TaobaoEticketMerchantMaDelayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
