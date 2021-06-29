package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发票扣减的接口 APIRequest
alibaba.einvoice.deduct.get

获取历史发票扣减量、每日发票扣减量的接口
*/
type AlibabaEinvoiceDeductGetRequest struct {
    model.Params

    // 税号
    payeeRegisterNo   string 

    // 业务日期
    bizDate   string 

    // 类型 1：所有 2：当日
    type   int64 

}

func NewAlibabaEinvoiceDeductGetRequest() *AlibabaEinvoiceDeductGetRequest{
    return &AlibabaEinvoiceDeductGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceDeductGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.deduct.get"
}

func (r AlibabaEinvoiceDeductGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceDeductGetRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceDeductGetRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

func (r *AlibabaEinvoiceDeductGetRequest) SetBizDate(bizDate string) error {
    r.bizDate = bizDate
    r.Set("biz_date", bizDate)
    return nil
}

func (r AlibabaEinvoiceDeductGetRequest) GetBizDate() string {
    return r.bizDate
}

func (r *AlibabaEinvoiceDeductGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r AlibabaEinvoiceDeductGetRequest) GetType() int64 {
    return r.type
}

