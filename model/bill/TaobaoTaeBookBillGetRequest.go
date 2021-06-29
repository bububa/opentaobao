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
    _bid   int64
    // 需要返回的字段:参见BookBill结构体
    _fields   []string
    // 虚拟账户流水编号
    _id   int64
    // 虚拟账户科目编号
    _accountId   int64
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
func (r *TaobaoTaeBookBillGetRequest) SetBid(_bid int64) error {
    r._bid = _bid
    r.Set("bid", _bid)
    return nil
}

// Bid Getter
func (r TaobaoTaeBookBillGetRequest) GetBid() int64 {
    return r._bid
}
// Fields Setter
// 需要返回的字段:参见BookBill结构体
func (r *TaobaoTaeBookBillGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTaeBookBillGetRequest) GetFields() []string {
    return r._fields
}
// Id Setter
// 虚拟账户流水编号
func (r *TaobaoTaeBookBillGetRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r TaobaoTaeBookBillGetRequest) GetId() int64 {
    return r._id
}
// AccountId Setter
// 虚拟账户科目编号
func (r *TaobaoTaeBookBillGetRequest) SetAccountId(_accountId int64) error {
    r._accountId = _accountId
    r.Set("account_id", _accountId)
    return nil
}

// AccountId Getter
func (r TaobaoTaeBookBillGetRequest) GetAccountId() int64 {
    return r._accountId
}
