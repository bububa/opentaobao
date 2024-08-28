package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaResend 电子凭证重发回调接口
// taobao.eticket.merchant.ma.resend
//
// 码商重发电子凭证回调接口
func TaobaoEticketMerchantMaResend(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaResendAPIRequest, resp *eticket.TaobaoEticketMerchantMaResendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
