package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔账单明细 API请求
taobao.tae.bill.get

查询单笔账单明细
*/
type TaobaoTaeBillGetRequest struct {
    model.Params
    // 账单编号
    bid   int64
    // 传入需要返回的字段
    fields   []string
    // 账单编号
    id   int64
    // 虚拟账户科目编号
    accountId   int64
}

// 初始化TaobaoTaeBillGetRequest对象
func NewTaobaoTaeBillGetRequest() *TaobaoTaeBillGetRequest{
    return &TaobaoTaeBillGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaeBillGetRequest) GetApiMethodName() string {
    return "taobao.tae.bill.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaeBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bid Setter
// 账单编号
func (r *TaobaoTaeBillGetRequest) SetBid(bid int64) error {
    r.bid = bid
    r.Set("bid", bid)
    return nil
}

// Bid Getter
func (r TaobaoTaeBillGetRequest) GetBid() int64 {
    return r.bid
}
// Fields Setter
// 传入需要返回的字段
func (r *TaobaoTaeBillGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TaobaoTaeBillGetRequest) GetFields() []string {
    return r.fields
}
// Id Setter
// 账单编号
func (r *TaobaoTaeBillGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

// Id Getter
func (r TaobaoTaeBillGetRequest) GetId() int64 {
    return r.id
}
// AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBillGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

// AccountId Getter
func (r TaobaoTaeBillGetRequest) GetAccountId() int64 {
    return r.accountId
}
