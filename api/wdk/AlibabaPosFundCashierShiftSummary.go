package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaPosFundCashierShiftSummary 收银换班数据同步
// alibaba.pos.fund.cashier.shift.summary
//
// 收银换班数据同步，将每天收银换班的数据回流给商家。
func AlibabaPosFundCashierShiftSummary(clt *core.SDKClient, req *wdk.AlibabaPosFundCashierShiftSummaryAPIRequest, resp *wdk.AlibabaPosFundCashierShiftSummaryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
