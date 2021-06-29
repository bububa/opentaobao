package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
极速开票开票请求回传 API请求
taobao.xhotel.fastinvoice.request

极速开票开票请求回传,用于记录航信开票请求数据
*/
type TaobaoXhotelFastinvoiceRequestRequest struct {
    model.Params
    // 无
    _invoiceInfoParam   *InvoiceInfoParam
}

// 初始化TaobaoXhotelFastinvoiceRequestRequest对象
func NewTaobaoXhotelFastinvoiceRequestRequest() *TaobaoXhotelFastinvoiceRequestRequest{
    return &TaobaoXhotelFastinvoiceRequestRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelFastinvoiceRequestRequest) GetApiMethodName() string {
    return "taobao.xhotel.fastinvoice.request"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelFastinvoiceRequestRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceInfoParam Setter
// 无
func (r *TaobaoXhotelFastinvoiceRequestRequest) SetInvoiceInfoParam(_invoiceInfoParam *InvoiceInfoParam) error {
    r._invoiceInfoParam = _invoiceInfoParam
    r.Set("invoice_info_param", _invoiceInfoParam)
    return nil
}

// InvoiceInfoParam Getter
func (r TaobaoXhotelFastinvoiceRequestRequest) GetInvoiceInfoParam() *InvoiceInfoParam {
    return r._invoiceInfoParam
}
