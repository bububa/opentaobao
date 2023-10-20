package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaposfundcashiershiftsummaryAPIRequest 收银换班数据同步 API请求
// alibaba.pos.fund.cashier.shift.summary
//
// 收银换班数据同步，将每天收银换班的数据回流给商家。
type AlibabaposfundcashiershiftsummaryAPIRequest struct {
	model.Params
	// 请求参数
	_cashierShiftFundRequest *CashierShiftFundRequest
}

// NewAlibabaposfundcashiershiftsummaryRequest 初始化AlibabaposfundcashiershiftsummaryAPIRequest对象
func NewAlibabaposfundcashiershiftsummaryRequest() *AlibabaposfundcashiershiftsummaryAPIRequest {
	return &AlibabaposfundcashiershiftsummaryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaposfundcashiershiftsummaryAPIRequest) GetApiMethodName() string {
	return "alibaba.pos.fund.cashier.shift.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaposfundcashiershiftsummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaposfundcashiershiftsummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCashierShiftFundRequest is CashierShiftFundRequest Setter
// 请求参数
func (r *AlibabaposfundcashiershiftsummaryAPIRequest) SetCashierShiftFundRequest(_cashierShiftFundRequest *CashierShiftFundRequest) error {
	r._cashierShiftFundRequest = _cashierShiftFundRequest
	r.Set("cashier_shift_fund_request", _cashierShiftFundRequest)
	return nil
}

// GetCashierShiftFundRequest CashierShiftFundRequest Getter
func (r AlibabaposfundcashiershiftsummaryAPIRequest) GetCashierShiftFundRequest() *CashierShiftFundRequest {
	return r._cashierShiftFundRequest
}
