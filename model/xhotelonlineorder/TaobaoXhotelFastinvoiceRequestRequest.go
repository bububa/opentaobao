package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
极速开票开票请求回传 APIRequest
taobao.xhotel.fastinvoice.request

极速开票开票请求回传,用于记录航信开票请求数据
*/
type TaobaoXhotelFastinvoiceRequestRequest struct {
    model.Params

    // 无
    invoiceInfoParam   *InvoiceInfoParam 

}

func NewTaobaoXhotelFastinvoiceRequestRequest() *TaobaoXhotelFastinvoiceRequestRequest{
    return &TaobaoXhotelFastinvoiceRequestRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelFastinvoiceRequestRequest) GetApiMethodName() string {
    return "taobao.xhotel.fastinvoice.request"
}

func (r TaobaoXhotelFastinvoiceRequestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelFastinvoiceRequestRequest) SetInvoiceInfoParam(invoiceInfoParam *InvoiceInfoParam) error {
    r.invoiceInfoParam = invoiceInfoParam
    r.Set("invoice_info_param", invoiceInfoParam)
    return nil
}

func (r TaobaoXhotelFastinvoiceRequestRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
    return r.invoiceInfoParam
}

