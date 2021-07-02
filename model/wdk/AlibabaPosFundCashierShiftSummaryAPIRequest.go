package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPosFundCashierShiftSummaryAPIRequest 收银换班数据同步 API请求
// alibaba.pos.fund.cashier.shift.summary
//
// 收银换班数据同步，将每天收银换班的数据回流给商家。
type AlibabaPosFundCashierShiftSummaryAPIRequest struct {
	model.Params
	// 请求参数
	_cashierShiftFundRequest *CashierShiftFundRequest
}

// NewAlibabaPosFundCashierShiftSummaryRequest 初始化AlibabaPosFundCashierShiftSummaryAPIRequest对象
func NewAlibabaPosFundCashierShiftSummaryRequest() *AlibabaPosFundCashierShiftSummaryAPIRequest {
	return &AlibabaPosFundCashierShiftSummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetApiMethodName() string {
	return "alibaba.pos.fund.cashier.shift.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CashierShiftFundRequest Setter
// 请求参数
func (r *AlibabaPosFundCashierShiftSummaryAPIRequest) SetCashierShiftFundRequest(_cashierShiftFundRequest *CashierShiftFundRequest) error {
	r._cashierShiftFundRequest = _cashierShiftFundRequest
	r.Set("cashier_shift_fund_request", _cashierShiftFundRequest)
	return nil
}

// Get CashierShiftFundRequest Getter
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetCashierShiftFundRequest() *CashierShiftFundRequest {
	return r._cashierShiftFundRequest
}
