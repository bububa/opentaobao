package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
极速开票开票请求完成 APIRequest
taobao.xhotel.fastinvoice.complete

极速开票开票请求回传,用于更新航信开票请求数据
*/
type TaobaoXhotelFastinvoiceCompleteRequest struct {
    model.Params

    // 无
    invoiceInfoParam   *InvoiceInfoParam 

}

func NewTaobaoXhotelFastinvoiceCompleteRequest() *TaobaoXhotelFastinvoiceCompleteRequest{
    return &TaobaoXhotelFastinvoiceCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelFastinvoiceCompleteRequest) GetApiMethodName() string {
    return "taobao.xhotel.fastinvoice.complete"
}

func (r TaobaoXhotelFastinvoiceCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelFastinvoiceCompleteRequest) SetInvoiceInfoParam(invoiceInfoParam *InvoiceInfoParam) error {
    r.invoiceInfoParam = invoiceInfoParam
    r.Set("invoice_info_param", invoiceInfoParam)
    return nil
}

func (r TaobaoXhotelFastinvoiceCompleteRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
    return r.invoiceInfoParam
}

