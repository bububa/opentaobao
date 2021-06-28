package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询费用科目信息(限自研商家) APIRequest
taobao.bill.accounts.get

查询费用账户信息
*/
type TaobaoBillAccountsGetRequest struct {
    model.Params

    // 需要返回的字段
    fields   []string 

    // 需要获取的科目ID
    aids   []int64 

}

func NewTaobaoBillAccountsGetRequest() *TaobaoBillAccountsGetRequest{
    return &TaobaoBillAccountsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBillAccountsGetRequest) GetApiMethodName() string {
    return "taobao.bill.accounts.get"
}

func (r TaobaoBillAccountsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBillAccountsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoBillAccountsGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoBillAccountsGetRequest) SetAids(aids []int64) error {
    r.aids = aids
    r.Set("aids", aids)
    return nil
}

func (r TaobaoBillAccountsGetRequest) GetAids() []int64 {
    return r.aids
}

