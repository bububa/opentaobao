package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillTicketPaymentQuery 获取影院票务账单-支付订单
// taobao.film.tfavatar.bill.ticket.payment.query
//
// 获取影院票务账单-支付订单
func TaobaoFilmTfavatarBillTicketPaymentQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillTicketPaymentQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
