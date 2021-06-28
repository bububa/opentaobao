package bill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询费用科目信息 APIRequest
taobao.tae.accounts.get

tae查询费用科目信息
*/
type TaobaoTaeAccountsGetRequest struct {
    model.Params

    // 需要返回的字段
    fields   []string 

    // 需要获取的科目ID
    aids   []int64 

}

func NewTaobaoTaeAccountsGetRequest() *TaobaoTaeAccountsGetRequest{
    return &TaobaoTaeAccountsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTaeAccountsGetRequest) GetApiMethodName() string {
    return "taobao.tae.accounts.get"
}

func (r TaobaoTaeAccountsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTaeAccountsGetRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TaobaoTaeAccountsGetRequest) GetFields() []string {
    return r.fields
}

func (r *TaobaoTaeAccountsGetRequest) SetAids(aids []int64) error {
    r.aids = aids
    r.Set("aids", aids)
    return nil
}

func (r TaobaoTaeAccountsGetRequest) GetAids() []int64 {
    return r.aids
}

