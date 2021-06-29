package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询费用科目信息(限自研商家) API请求
taobao.bill.accounts.get

查询费用账户信息
*/
type TaobaoBillAccountsGetRequest struct {
    model.Params
    // 需要返回的字段
    _fields   []string
    // 需要获取的科目ID
    _aids   []int64
}

// 初始化TaobaoBillAccountsGetRequest对象
func NewTaobaoBillAccountsGetRequest() *TaobaoBillAccountsGetRequest{
    return &TaobaoBillAccountsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBillAccountsGetRequest) GetApiMethodName() string {
    return "taobao.bill.accounts.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBillAccountsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段
func (r *TaobaoBillAccountsGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoBillAccountsGetRequest) GetFields() []string {
    return r._fields
}
// Aids Setter
// 需要获取的科目ID
func (r *TaobaoBillAccountsGetRequest) SetAids(_aids []int64) error {
    r._aids = _aids
    r.Set("aids", _aids)
    return nil
}

// Aids Getter
func (r TaobaoBillAccountsGetRequest) GetAids() []int64 {
    return r._aids
}
