package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
开票量核对接口 API请求
alibaba.einvoice.amount.check

跟开票服务商核对历史开票量，用来对账
*/
type AlibabaEinvoiceAmountCheckRequest struct {
    model.Params
    // 税号
    _payeeRegisterNo   string
    // 开票日期开始时间
    _startDate   string
    // 开票日期结束时间
    _endDate   string
}

// 初始化AlibabaEinvoiceAmountCheckRequest对象
func NewAlibabaEinvoiceAmountCheckRequest() *AlibabaEinvoiceAmountCheckRequest{
    return &AlibabaEinvoiceAmountCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceAmountCheckRequest) GetApiMethodName() string {
    return "alibaba.einvoice.amount.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceAmountCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceAmountCheckRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceAmountCheckRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// StartDate Setter
// 开票日期开始时间
func (r *AlibabaEinvoiceAmountCheckRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaEinvoiceAmountCheckRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 开票日期结束时间
func (r *AlibabaEinvoiceAmountCheckRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaEinvoiceAmountCheckRequest) GetEndDate() string {
    return r._endDate
}
