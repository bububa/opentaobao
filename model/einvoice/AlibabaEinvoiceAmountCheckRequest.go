package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票量核对接口 APIRequest
alibaba.einvoice.amount.check

跟开票服务商核对历史开票量，用来对账
*/
type AlibabaEinvoiceAmountCheckRequest struct {
    model.Params

    // 税号
    payeeRegisterNo   string 

    // 开票日期开始时间
    startDate   string 

    // 开票日期结束时间
    endDate   string 

}

func NewAlibabaEinvoiceAmountCheckRequest() *AlibabaEinvoiceAmountCheckRequest{
    return &AlibabaEinvoiceAmountCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceAmountCheckRequest) GetApiMethodName() string {
    return "alibaba.einvoice.amount.check"
}

func (r AlibabaEinvoiceAmountCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceAmountCheckRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceAmountCheckRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceAmountCheckRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

func (r AlibabaEinvoiceAmountCheckRequest) GetStartDate() string {
    return r.startDate
}

func (r *AlibabaEinvoiceAmountCheckRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

func (r AlibabaEinvoiceAmountCheckRequest) GetEndDate() string {
    return r.endDate
}

