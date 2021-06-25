package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔虚拟账户明细 APIRequest
taobao.tae.book.bill.get

tae查询单笔虚拟账户明细
*/
type TaobaoTaeBookBillGetRequest struct {
    model.Params

    // 虚拟账户流水编号
    bid   int64 

    // 需要返回的字段:参见BookBill结构体
    fields   []String 

    // 虚拟账户流水编号
    id   int64 

    // 虚拟账户科目编号
    accountId   int64 

}

func NewTaobaoTaeBookBillGetRequest() *TaobaoTaeBookBillGetRequest{
    return &TaobaoTaeBookBillGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaeBookBillGetRequest) GetApiMethodName() string {
    return "taobao.tae.book.bill.get"
}

func (r TaobaoTaeBookBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaeBookBillGetRequest) SetBid(bid int64) error {
    r.bid = bid
    r.Set("bid", bid)
    return nil
}

func (r TaobaoTaeBookBillGetRequest) GetBid() int64 {
    return r.bid
}

func (r *TaobaoTaeBookBillGetRequest) SetFields(fields []String) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTaeBookBillGetRequest) GetFields() []String {
    return r.fields
}

func (r *TaobaoTaeBookBillGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoTaeBookBillGetRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoTaeBookBillGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r TaobaoTaeBookBillGetRequest) GetAccountId() int64 {
    return r.accountId
}

