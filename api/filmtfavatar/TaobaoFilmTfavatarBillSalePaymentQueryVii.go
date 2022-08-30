package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSalePaymentQueryVii 获取影院卖品账单--支付账单-V2版本(正逆分离)
// taobao.film.tfavatar.bill.sale.payment.query.vii
//
// 获取影院卖品账单--支付账单-V2版本(正逆分离)
func TaobaoFilmTfavatarBillSalePaymentQueryVii(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryViiAPIRequest, session string) (*filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponse, error) {
	var resp filmtfavatar.TaobaoFilmTfavatarBillSalePaymentQueryViiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
