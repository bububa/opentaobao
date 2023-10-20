package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSalePaymentQuery 获取影院卖品账单--支付账单
// taobao.film.tfavatar.bill.sale.payment.query
//
// 获取影院卖品账单--支付账单
func TaobaoFilmTfavatarBillSalePaymentQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
