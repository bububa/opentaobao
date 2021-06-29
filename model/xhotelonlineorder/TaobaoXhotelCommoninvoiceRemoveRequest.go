package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
常用发票信息删除接口 APIRequest
taobao.xhotel.commoninvoice.remove

常用发票信息删除接口
*/
type TaobaoXhotelCommoninvoiceRemoveRequest struct {
    model.Params

    // 发票id
    invoiceId   int64 

    // 用户名
    userNick   string 

}

func NewTaobaoXhotelCommoninvoiceRemoveRequest() *TaobaoXhotelCommoninvoiceRemoveRequest{
    return &TaobaoXhotelCommoninvoiceRemoveRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetApiMethodName() string {
    return "taobao.xhotel.commoninvoice.remove"
}

func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelCommoninvoiceRemoveRequest) SetInvoiceId(invoiceId int64) error {
    r.invoiceId = invoiceId
    r.Set("invoice_id", invoiceId)
    return nil
}

func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetInvoiceId() int64 {
    return r.invoiceId
}

func (r *TaobaoXhotelCommoninvoiceRemoveRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetUserNick() string {
    return r.userNick
}

