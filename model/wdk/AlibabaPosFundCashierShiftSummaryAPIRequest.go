package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPosFundCashierShiftSummaryAPIRequest) Reset() {
	r._cashierShiftFundRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetApiMethodName() string {
	return "alibaba.pos.fund.cashier.shift.summary"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCashierShiftFundRequest is CashierShiftFundRequest Setter
// 请求参数
func (r *AlibabaPosFundCashierShiftSummaryAPIRequest) SetCashierShiftFundRequest(_cashierShiftFundRequest *CashierShiftFundRequest) error {
	r._cashierShiftFundRequest = _cashierShiftFundRequest
	r.Set("cashier_shift_fund_request", _cashierShiftFundRequest)
	return nil
}

// GetCashierShiftFundRequest CashierShiftFundRequest Getter
func (r AlibabaPosFundCashierShiftSummaryAPIRequest) GetCashierShiftFundRequest() *CashierShiftFundRequest {
	return r._cashierShiftFundRequest
}

var poolAlibabaPosFundCashierShiftSummaryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPosFundCashierShiftSummaryRequest()
	},
}

// GetAlibabaPosFundCashierShiftSummaryRequest 从 sync.Pool 获取 AlibabaPosFundCashierShiftSummaryAPIRequest
func GetAlibabaPosFundCashierShiftSummaryAPIRequest() *AlibabaPosFundCashierShiftSummaryAPIRequest {
	return poolAlibabaPosFundCashierShiftSummaryAPIRequest.Get().(*AlibabaPosFundCashierShiftSummaryAPIRequest)
}

// ReleaseAlibabaPosFundCashierShiftSummaryAPIRequest 将 AlibabaPosFundCashierShiftSummaryAPIRequest 放入 sync.Pool
func ReleaseAlibabaPosFundCashierShiftSummaryAPIRequest(v *AlibabaPosFundCashierShiftSummaryAPIRequest) {
	v.Reset()
	poolAlibabaPosFundCashierShiftSummaryAPIRequest.Put(v)
}
