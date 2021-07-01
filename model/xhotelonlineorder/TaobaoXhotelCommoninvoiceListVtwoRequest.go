package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户常用发票信息查询接口 API请求
taobao.xhotel.commoninvoice.list.vtwo

获取用户常用发票信息接口
*/
type TaobaoXhotelCommoninvoiceListVtwoAPIRequest struct {
    model.Params
}

// 初始化TaobaoXhotelCommoninvoiceListVtwoAPIRequest对象
func NewTaobaoXhotelCommoninvoiceListVtwoRequest() *TaobaoXhotelCommoninvoiceListVtwoAPIRequest{
    return &TaobaoXhotelCommoninvoiceListVtwoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceListVtwoAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.commoninvoice.list.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceListVtwoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
