package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询费用科目信息 API请求
taobao.tae.accounts.get

tae查询费用科目信息
*/
type TaobaoTaeAccountsGetRequest struct {
    model.Params
    // 需要返回的字段
    _fields   []string
    // 需要获取的科目ID
    _aids   []int64
}

// 初始化TaobaoTaeAccountsGetRequest对象
func NewTaobaoTaeAccountsGetRequest() *TaobaoTaeAccountsGetRequest{
    return &TaobaoTaeAccountsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTaeAccountsGetRequest) GetApiMethodName() string {
    return "taobao.tae.accounts.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTaeAccountsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Fields Setter
// 需要返回的字段
func (r *TaobaoTaeAccountsGetRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TaobaoTaeAccountsGetRequest) GetFields() []string {
    return r._fields
}
// Aids Setter
// 需要获取的科目ID
func (r *TaobaoTaeAccountsGetRequest) SetAids(_aids []int64) error {
    r._aids = _aids
    r.Set("aids", _aids)
    return nil
}

// Aids Getter
func (r TaobaoTaeAccountsGetRequest) GetAids() []int64 {
    return r._aids
}
