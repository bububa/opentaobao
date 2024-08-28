package filmtfavatar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillTicketPaymentQuery 获取影院票务账单-支付订单
// taobao.film.tfavatar.bill.ticket.payment.query
//
// 获取影院票务账单-支付订单
func TaobaoFilmTfavatarBillTicketPaymentQuery(ctx context.Context, clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
