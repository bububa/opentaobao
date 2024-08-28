package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoCardExpandcardQuery 购物金卡查询
// taobao.card.expandcard.query
//
// 购物金充值信息查询接口，会返回余额等信息。
func TaobaoCardExpandcardQuery(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoCardExpandcardQueryAPIRequest, resp *promotion.TaobaoCardExpandcardQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
