package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一批账号ID查询账号列表 APIRequest
alibaba.mozi.buc.account.list.accountids

根据一批账号ID查询账号列表
*/
type AlibabaMoziBucAccountListAccountidsRequest struct {
    model.Params

    // 请求参数
    listAccountIds   *ListAccountsByAccountIdsRequest 

}

func NewAlibabaMoziBucAccountListAccountidsRequest() *AlibabaMoziBucAccountListAccountidsRequest{
    return &AlibabaMoziBucAccountListAccountidsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziBucAccountListAccountidsRequest) GetApiMethodName() string {
    return "alibaba.mozi.buc.account.list.accountids"
}

func (r AlibabaMoziBucAccountListAccountidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziBucAccountListAccountidsRequest) SetListAccountIds(listAccountIds *ListAccountsByAccountIdsRequest) error {
    r.listAccountIds = listAccountIds
    r.Set("list_account_ids", listAccountIds)
    return nil
}

func (r AlibabaMoziBucAccountListAccountidsRequest) GetListAccountIds() *ListAccountsByAccountIdsRequest {
    return r.listAccountIds
}

