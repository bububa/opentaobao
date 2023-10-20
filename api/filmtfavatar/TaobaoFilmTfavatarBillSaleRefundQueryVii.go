package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSaleRefundQueryVii 获取影院卖品账单--退款账单-V2版本(正逆分离)
// taobao.film.tfavatar.bill.sale.refund.query.vii
//
// 获取影院卖品账单--退款账单-V2版本(正逆分离)
func TaobaoFilmTfavatarBillSaleRefundQueryVii(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryViiAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryViiAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
