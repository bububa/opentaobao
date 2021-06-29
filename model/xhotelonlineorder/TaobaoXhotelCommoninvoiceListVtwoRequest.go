package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户常用发票信息查询接口 APIRequest
taobao.xhotel.commoninvoice.list.vtwo

获取用户常用发票信息接口
*/
type TaobaoXhotelCommoninvoiceListVtwoRequest struct {
    model.Params

}

func NewTaobaoXhotelCommoninvoiceListVtwoRequest() *TaobaoXhotelCommoninvoiceListVtwoRequest{
    return &TaobaoXhotelCommoninvoiceListVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelCommoninvoiceListVtwoRequest) GetApiMethodName() string {
    return "taobao.xhotel.commoninvoice.list.vtwo"
}

func (r TaobaoXhotelCommoninvoiceListVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


