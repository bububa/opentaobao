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
    _serialNo   string
    // 税号
    _payeeRegisterNo   string
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
func (r *AlibabaEinvoiceClosereqRequest) SetSerialNo(_serialNo string) error {
    r._serialNo = _serialNo
    r.Set("serial_no", _serialNo)
    return nil
}

// SerialNo Getter
func (r AlibabaEinvoiceClosereqRequest) GetSerialNo() string {
    return r._serialNo
}
// PayeeRegisterNo Setter
// 税号
func (r *AlibabaEinvoiceClosereqRequest) SetPayeeRegisterNo(_payeeRegisterNo string) error {
    r._payeeRegisterNo = _payeeRegisterNo
    r.Set("payee_register_no", _payeeRegisterNo)
    return nil
}

// PayeeRegisterNo Getter
func (r AlibabaEinvoiceClosereqRequest) GetPayeeRegisterNo() string {
    return r._payeeRegisterNo
}
