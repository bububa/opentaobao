package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
常用发票信息更新接口 APIRequest
taobao.xhotel.commoninvoice.update

常用发票信息更新接口(根据用户id,发票抬头和发票属性或发票id进行更新,没有则添加)
*/
type TaobaoXhotelCommoninvoiceUpdateRequest struct {
    model.Params

    // 无
    commonInvoiceInfoParam   *CommonInvoiceInfo 

}

func NewTaobaoXhotelCommoninvoiceUpdateRequest() *TaobaoXhotelCommoninvoiceUpdateRequest{
    return &TaobaoXhotelCommoninvoiceUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelCommoninvoiceUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.commoninvoice.update"
}

func (r TaobaoXhotelCommoninvoiceUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelCommoninvoiceUpdateRequest) SetCommonInvoiceInfoParam(commonInvoiceInfoParam *CommonInvoiceInfo) error {
    r.commonInvoiceInfoParam = commonInvoiceInfoParam
    r.Set("common_invoice_info_param", commonInvoiceInfoParam)
    return nil
}

func (r TaobaoXhotelCommoninvoiceUpdateRequest) GetCommonInvoiceInfoParam() *CommonInvoiceInfo {
    return r.commonInvoiceInfoParam
}

