package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoServindustryFinanceGeexOrderLoan 即科放款信息回调api
// taobao.servindustry.finance.geex.order.loan
//
// 即科放款信息api
func TaobaoServindustryFinanceGeexOrderLoan(clt *core.SDKClient, req *trade.TaobaoServindustryFinanceGeexOrderLoanAPIRequest, resp *trade.TaobaoServindustryFinanceGeexOrderLoanAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
