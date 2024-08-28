package filmtfavatar

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/filmtfavatar"
)

// TaobaoFilmTfavatarBillSalePrintQuery 获取影院卖品账单-核销账单
// taobao.film.tfavatar.bill.sale.print.query
//
// 获取影院卖品账单-核销账单
// 返回值data属于加密字段, 并非大字段.
func TaobaoFilmTfavatarBillSalePrintQuery(ctx context.Context, clt *core.SDKClient, req *filmtfavatar.TaobaoFilmTfavatarBillSalePrintQueryAPIRequest, resp *filmtfavatar.TaobaoFilmTfavatarBillSalePrintQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
