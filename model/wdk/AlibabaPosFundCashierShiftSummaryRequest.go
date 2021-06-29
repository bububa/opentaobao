package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
收银换班数据同步 API请求
alibaba.pos.fund.cashier.shift.summary

收银换班数据同步，将每天收银换班的数据回流给商家。
*/
type AlibabaPosFundCashierShiftSummaryRequest struct {
    model.Params
    // 请求参数
    _cashierShiftFundRequest   *CashierShiftFundRequest
}

// 初始化AlibabaPosFundCashierShiftSummaryRequest对象
func NewAlibabaPosFundCashierShiftSummaryRequest() *AlibabaPosFundCashierShiftSummaryRequest{
    return &AlibabaPosFundCashierShiftSummaryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPosFundCashierShiftSummaryRequest) GetApiMethodName() string {
    return "alibaba.pos.fund.cashier.shift.summary"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPosFundCashierShiftSummaryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CashierShiftFundRequest Setter
// 请求参数
func (r *AlibabaPosFundCashierShiftSummaryRequest) SetCashierShiftFundRequest(_cashierShiftFundRequest *CashierShiftFundRequest) error {
    r._cashierShiftFundRequest = _cashierShiftFundRequest
    r.Set("cashier_shift_fund_request", _cashierShiftFundRequest)
    return nil
}

// CashierShiftFundRequest Getter
func (r AlibabaPosFundCashierShiftSummaryRequest) GetCashierShiftFundRequest() *CashierShiftFundRequest {
    return r._cashierShiftFundRequest
}
