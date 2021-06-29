package einvoice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家开票申请单状态回传 APIRequest
alibaba.einvoice.invoiceapply.update

开票服务商更新商家开票申请单状态
*/
type AlibabaEinvoiceInvoiceapplyUpdateRequest struct {
    model.Params

    // 申请单id
    applyId   string 

    // 申请单状态：1：待确认，2：开票中，3：拒绝开票，4：发票已发出，0：完成开票
    status   int64 

    // 扩展信息,目前用于回传文本及物流消息
    exInfo   string 

}

func NewAlibabaEinvoiceInvoiceapplyUpdateRequest() *AlibabaEinvoiceInvoiceapplyUpdateRequest{
    return &AlibabaEinvoiceInvoiceapplyUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaEinvoiceInvoiceapplyUpdateRequest) GetApiMethodName() string {
    return "alibaba.einvoice.invoiceapply.update"
}

func (r AlibabaEinvoiceInvoiceapplyUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaEinvoiceInvoiceapplyUpdateRequest) SetApplyId(applyId string) error {
    r.applyId = applyId
    r.Set("apply_id", applyId)
    return nil
}

func (r AlibabaEinvoiceInvoiceapplyUpdateRequest) GetApplyId() string {
    return r.applyId
}

func (r *AlibabaEinvoiceInvoiceapplyUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r AlibabaEinvoiceInvoiceapplyUpdateRequest) GetStatus() int64 {
    return r.status
}

func (r *AlibabaEinvoiceInvoiceapplyUpdateRequest) SetExInfo(exInfo string) error {
    r.exInfo = exInfo
    r.Set("ex_info", exInfo)
    return nil
}

func (r AlibabaEinvoiceInvoiceapplyUpdateRequest) GetExInfo() string {
    return r.exInfo
}

