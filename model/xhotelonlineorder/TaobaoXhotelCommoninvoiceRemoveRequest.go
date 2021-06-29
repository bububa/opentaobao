package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
常用发票信息删除接口 API请求
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

// 初始化TaobaoXhotelCommoninvoiceRemoveRequest对象
func NewTaobaoXhotelCommoninvoiceRemoveRequest() *TaobaoXhotelCommoninvoiceRemoveRequest{
    return &TaobaoXhotelCommoninvoiceRemoveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetApiMethodName() string {
    return "taobao.xhotel.commoninvoice.remove"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InvoiceId Setter
// 发票id
func (r *TaobaoXhotelCommoninvoiceRemoveRequest) SetInvoiceId(invoiceId int64) error {
    r.invoiceId = invoiceId
    r.Set("invoice_id", invoiceId)
    return nil
}

// InvoiceId Getter
func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetInvoiceId() int64 {
    return r.invoiceId
}
// UserNick Setter
// 用户名
func (r *TaobaoXhotelCommoninvoiceRemoveRequest) SetUserNick(userNick string) error {
    r.userNick = userNick
    r.Set("user_nick", userNick)
    return nil
}

// UserNick Getter
func (r TaobaoXhotelCommoninvoiceRemoveRequest) GetUserNick() string {
    return r.userNick
}
