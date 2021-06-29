package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的开票申请 API请求
alibaba.einvoice.invoiceapply.get

开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
*/
type AlibabaEinvoiceInvoiceapplyGetRequest struct {
    model.Params
    // 开票申请id
    _applyId   string
}

// 初始化AlibabaEinvoiceInvoiceapplyGetRequest对象
func NewAlibabaEinvoiceInvoiceapplyGetRequest() *AlibabaEinvoiceInvoiceapplyGetRequest{
    return &AlibabaEinvoiceInvoiceapplyGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaEinvoiceInvoiceapplyGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.invoiceapply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaEinvoiceInvoiceapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ApplyId Setter
// 开票申请id
func (r *AlibabaEinvoiceInvoiceapplyGetRequest) SetApplyId(_applyId string) error {
    r._applyId = _applyId
    r.Set("apply_id", _applyId)
    return nil
}

// ApplyId Getter
func (r AlibabaEinvoiceInvoiceapplyGetRequest) GetApplyId() string {
    return r._applyId
}
