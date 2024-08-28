package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoServindustryFinanceGeexOrderLoan 即科放款信息回调api
// taobao.servindustry.finance.geex.order.loan
//
// 即科放款信息api
func TaobaoServindustryFinanceGeexOrderLoan(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoServindustryFinanceGeexOrderLoanAPIRequest, resp *trade.TaobaoServindustryFinanceGeexOrderLoanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
