package filmtfavatar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSaleRefundQuery 获取影院卖品账单--退款账单
// taobao.film.tfavatar.bill.sale.refund.query
//
// 获取影院卖品账单--退款账单
func TaobaoFilmTfavatarBillSaleRefundQuery(ctx context.Context, clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillSaleRefundQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
