package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseExistinghomeRentTradeBindItem 交易绑定商品
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
func AlibabaAlihouseExistinghomeRentTradeBindItem(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseExistinghomeRentTradeBindItemAPIRequest, resp *alihouse.AlibabaAlihouseExistinghomeRentTradeBindItemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
