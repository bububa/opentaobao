package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaConsume 电子凭证核销接口
// taobao.eticket.merchant.ma.consume
//
// 电子凭证核销接口
func TaobaoEticketMerchantMaConsume(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaConsumeAPIRequest, resp *eticket.TaobaoEticketMerchantMaConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
