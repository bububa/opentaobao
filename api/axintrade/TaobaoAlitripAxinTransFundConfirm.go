package axintrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// TaobaoAlitripAxinTransFundConfirm 确认资金单
// taobao.alitrip.axin.trans.fund.confirm
//
// 通过外部订单号进行资金结算
func TaobaoAlitripAxinTransFundConfirm(ctx context.Context, clt *core.SDKClient, req *axintrade.TaobaoAlitripAxinTransFundConfirmAPIRequest, resp *axintrade.TaobaoAlitripAxinTransFundConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
