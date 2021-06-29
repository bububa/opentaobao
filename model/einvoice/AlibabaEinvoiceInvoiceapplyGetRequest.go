package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取商家的开票申请 APIRequest
alibaba.einvoice.invoiceapply.get

开票服务商接收到商家发起的开票申请消息后，调用此接口拉取商家详细的开票申请内容
*/
type AlibabaEinvoiceInvoiceapplyGetRequest struct {
    model.Params

    // 开票申请id
    applyId   string 

}

func NewAlibabaEinvoiceInvoiceapplyGetRequest() *AlibabaEinvoiceInvoiceapplyGetRequest{
    return &AlibabaEinvoiceInvoiceapplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceInvoiceapplyGetRequest) GetApiMethodName() string {
    return "alibaba.einvoice.invoiceapply.get"
}

func (r AlibabaEinvoiceInvoiceapplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceInvoiceapplyGetRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlibabaEinvoiceInvoiceapplyGetRequest) GetApplyId() string {
    return r.applyId
}

