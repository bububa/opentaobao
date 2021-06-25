package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
OpenAccount账号信息查询 APIRequest
taobao.open.account.list

OpenAccount账号信息查询
*/
type TaobaoOpenAccountListRequest struct {
    model.Params

    // Open Account的id列表, 每次最多查询 20 个帐户
    openAccountIds   []Number 

    // ISV自己账号的id列表，isvAccountId和openAccountId二选一必填, 每次最多查询 20 个帐户
    isvAccountIds   []String 

}

func NewTaobaoOpenAccountListRequest() *TaobaoOpenAccountListRequest{
    return &TaobaoOpenAccountListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenAccountListRequest) GetApiMethodName() string {
    return "taobao.open.account.list"
}

func (r TaobaoOpenAccountListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenAccountListRequest) SetOpenAccountIds(openAccountIds []Number) error {
    r.openAccountIds = openAccountIds
    r.Set("open_account_ids", openAccountIds)
    return nil
}

func (r TaobaoOpenAccountListRequest) GetOpenAccountIds() []Number {
    return r.openAccountIds
}

func (r *TaobaoOpenAccountListRequest) SetIsvAccountIds(isvAccountIds []String) error {
    r.isvAccountIds = isvAccountIds
    r.Set("isv_account_ids", isvAccountIds)
    return nil
}

func (r TaobaoOpenAccountListRequest) GetIsvAccountIds() []String {
    return r.isvAccountIds
}

