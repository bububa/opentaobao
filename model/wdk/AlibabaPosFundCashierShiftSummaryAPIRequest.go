package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPosFundCashierShiftSummaryAPIRequest
收银换班数据同步 API请求
alibaba.pos.fund.cashier.shift.summary

收银换班数据同步，将每天收银换班的数据回流给商家。 */
type AlibabaPosFundCashierShiftSummaryAPIRequest struct {
	model.Params
	// 请求参数
	_cashierShiftFundRequest *CashierShiftFundRequest
}

// New
