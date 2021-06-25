package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount删除数据 APIRequest
taobao.open.account.delete

OpenAccount删除数据
*/
type TaobaoOpenAccountDeleteRequest struct {
    model.Params

    // Open Account的id列表
    openAccountIds   []Number 

    // ISV自己账号的id列表
    isvAccountIds   []String 

}

func NewTaobaoOpenAccountDeleteRequest() *TaobaoOpenAccountDeleteRequest{
    return &TaobaoOpenAccountDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountDeleteRequest) GetApiMethodName() string {
    return "taobao.open.account.delete"
}

func (r TaobaoOpenAccountDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountDeleteRequest) SetOpenAccountIds(openAccountIds []Number) error {
    r.openAccountIds = openAccountIds
    r.Set("open_account_ids", openAccountIds)
    return nil
}

func (r TaobaoOpenAccountDeleteRequest) GetOpenAccountIds() []Number {
    return r.openAccountIds
}

func (r *TaobaoOpenAccountDeleteRequest) SetIsvAccountIds(isvAccountIds []String) error {
    r.isvAccountIds = isvAccountIds
    r.Set("isv_account_ids", isvAccountIds)
    return nil
}

func (r TaobaoOpenAccountDeleteRequest) GetIsvAccountIds() []String {
    return r.isvAccountIds
}

