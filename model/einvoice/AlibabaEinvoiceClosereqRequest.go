package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭开票失败请求（失败列表可重试） APIRequest
alibaba.einvoice.closereq

关闭失败开票请求，避免造成重复开票
*/
type AlibabaEinvoiceClosereqRequest struct {
    model.Params

    // 流水号
    serialNo   string 

    // 税号
    payeeRegisterNo   string 

}

func NewAlibabaEinvoiceClosereqRequest() *AlibabaEinvoiceClosereqRequest{
    return &AlibabaEinvoiceClosereqRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceClosereqRequest) GetApiMethodName() string {
    return "alibaba.einvoice.closereq"
}

func (r AlibabaEinvoiceClosereqRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceClosereqRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

func (r AlibabaEinvoiceClosereqRequest) GetSerialNo() string {
    return r.serialNo
}

func (r *AlibabaEinvoiceClosereqRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

func (r AlibabaEinvoiceClosereqRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}

