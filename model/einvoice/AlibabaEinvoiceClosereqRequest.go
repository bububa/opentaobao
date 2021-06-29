package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关闭开票失败请求（失败列表可重试） API请求
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

// 初始化AlibabaEinvoiceClosereqRequest对象
func NewAlibabaEinvoiceClosereqRequest() *AlibabaEinvoiceClosereqRequest{
    return &AlibabaEinvoiceClosereqRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceClosereqRequest) GetApiMethodName() string {
    return "alibaba.einvoice.closereq"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceClosereqRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SerialNo Setter
// 流水号
func (r *AlibabaEinvoiceClosereqRequest) SetSerialNo(serialNo string) error {
    r.serialNo = serialNo
    r.Set("serial_no", serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceClosereqRequest) GetSerialNo() string {
    return r.serialNo
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceClosereqRequest) SetPayeeRegisterNo(payeeRegisterNo string) error {
    r.payeeRegisterNo = payeeRegisterNo
    r.Set("payee_register_no", payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceClosereqRequest) GetPayeeRegisterNo() string {
    return r.payeeRegisterNo
}
