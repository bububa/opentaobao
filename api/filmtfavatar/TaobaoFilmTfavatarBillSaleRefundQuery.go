package filmtfavatar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSaleRefundQuery 获取影院卖品账单--退款账单
// taobao.film.tfavatar.bill.sale.refund.query
//
// 获取影院卖品账单--退款账单
func TaobaoFilmTfavatarBillSaleRefundQuery(clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
