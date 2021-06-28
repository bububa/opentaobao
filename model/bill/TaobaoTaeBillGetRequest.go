package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔账单明细 APIRequest
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

func NewTaobaoTaeBillGetRequest() *TaobaoTaeBillGetRequest{
    return &TaobaoTaeBillGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaeBillGetRequest) GetApiMethodName() string {
    return "taobao.tae.bill.get"
}

func (r TaobaoTaeBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaeBillGetRequest) SetBid(bid int64) error {
    r.bid = bid
    r.Set("bid", bid)
    return nil
}

func (r TaobaoTaeBillGetRequest) GetBid() int64 {
    return r.bid
}

func (r *TaobaoTaeBillGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTaeBillGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoTaeBillGetRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TaobaoTaeBillGetRequest) GetId() int64 {
    return r.id
}

func (r *TaobaoTaeBillGetRequest) SetAccountId(accountId int64) error {
    r.accountId = accountId
    r.Set("account_id", accountId)
    return nil
}

func (r TaobaoTaeBillGetRequest) GetAccountId() int64 {
    return r.accountId
}

