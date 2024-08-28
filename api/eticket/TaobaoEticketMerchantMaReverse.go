package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaReverse 电子凭证冲正接口
// taobao.eticket.merchant.ma.reverse
//
// 电子凭证平台冲正接口
func TaobaoEticketMerchantMaReverse(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaReverseAPIRequest, resp *eticket.TaobaoEticketMerchantMaReverseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
