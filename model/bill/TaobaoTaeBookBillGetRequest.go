package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔虚拟账户明细 API请求
taobao.tae.book.bill.get

tae查询单笔虚拟账户明细
*/
type TaobaoTaeBookBillGetRequest struct {
    model.Params
    // 虚拟账户流水编号
    bid   int64
    // 需要返回的字段:参见BookBill结构体
    fields   []string
    // 虚拟账户流水编号
    id   int64
    // 虚拟账户科目编号
    accountId   int64
}

// 初始化TaobaoTaeBookBillGetRequest对象
func NewTaobaoTaeBookBillGetRequest() *TaobaoTaeBookBillGetRequest{
    return &TaobaoTaeBookBillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaeBookBillGetRequest) GetApiMethodName() string {
    return "taobao.tae.book.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaeBookBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bid Setter
// 虚拟账户流水编号
func (r *TaobaoTaeBookBillGetRequest) SetBid(bid int64) error {
    r.bid = bid
    r.Set("bid", bid)
    return nil
}

// Bid Getter
func (r TaobaoTaeBookBillGetRequest) GetBid() int64 {
    return r.bid
}
// Fields Setter
// 需要返回的字段:参见BookBill结构体
func (r *TaobaoTaeBookBillGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTaeBookBillGetRequest) GetFields() []string {
    return r.fields
}
// Id Setter
// 虚拟账户流水编号
func (r *TaobaoTaeBookBillGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoTaeBookBillGetRequest) GetId() int64 {
    return r.id
}
// AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBookBillGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r TaobaoTaeBookBillGetRequest) GetAccountId() int64 {
    return r.accountId
}
