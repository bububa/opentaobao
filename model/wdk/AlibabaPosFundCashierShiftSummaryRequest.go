package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
收银换班数据同步 APIRequest
alibaba.pos.fund.cashier.shift.summary

收银换班数据同步，将每天收银换班的数据回流给商家。
*/
type AlibabaPosFundCashierShiftSummaryRequest struct {
    model.Params

    // 请求参数
    cashierShiftFundRequest   *CashierShiftFundRequest 

}

func NewAlibabaPosFundCashierShiftSummaryRequest() *AlibabaPosFundCashierShiftSummaryRequest{
    return &AlibabaPosFundCashierShiftSummaryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaPosFundCashierShiftSummaryRequest) GetApiMethodName() string {
    return "alibaba.pos.fund.cashier.shift.summary"
}

func (r AlibabaPosFundCashierShiftSummaryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaPosFundCashierShiftSummaryRequest) SetCashierShiftFundRequest(cashierShiftFundRequest *CashierShiftFundRequest) error {
    r.cashierShiftFundRequest = cashierShiftFundRequest
    r.Set("cashier_shift_fund_request", cashierShiftFundRequest)
    return nil
}

func (r AlibabaPosFundCashierShiftSummaryRequest) GetCashierShiftFundRequest() *CashierShiftFundRequest {
    return r.cashierShiftFundRequest
}

