package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票扣减的接口 API请求
alibaba.einvoice.deduct.get

获取历史发票扣减量、每日发票扣减量的接口
*/
type AlibabaEinvoiceDeductGetRequest struct {
    model.Params
    // 税号
    _payeeRegisterNo   string
    // 业务日期
    _bizDate   string
    // 类型 1：所有 2：当日
    _type   int64
}

// 初始化AlibabaEinvoiceDeductGetRequest对象
func NewAlibabaEinvoiceDeductGetRequest() *AlibabaEinvoiceDeductGetRequest{
    return &AlibabaEinvoiceDeductGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceDeductGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.deduct.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceDeductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceDeductGetRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceDeductGetRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
// BizDate Setter
// 业务日期
func (r *AlibabaEinvoiceDeductGetRequest) SetBizDate(_bizDate string) error {
    r._bizDate = _bizDate
    r.Set("biz_date", _bizDate)
    return nil
}

// BizDate Getter
func (r AlibabaEinvoiceDeductGetRequest) GetBizDate() string {
    return r._bizDate
}
// Type Setter
// 类型 1：所有 2：当日
func (r *AlibabaEinvoiceDeductGetRequest) SetType(_type int64) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaEinvoiceDeductGetRequest) GetType() int64 {
    return r._type
}
