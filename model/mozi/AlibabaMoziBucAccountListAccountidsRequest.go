package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一批账号ID查询账号列表 API请求
alibaba.mozi.buc.account.list.accountids

根据一批账号ID查询账号列表
*/
type AlibabaMoziBucAccountListAccountidsRequest struct {
    model.Params
    // 请求参数
    _listAccountIds   *ListAccountsByAccountIdsRequest
}

// 初始化AlibabaMoziBucAccountListAccountidsRequest对象
func NewAlibabaMoziBucAccountListAccountidsRequest() *AlibabaMoziBucAccountListAccountidsRequest{
    return &AlibabaMoziBucAccountListAccountidsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMoziBucAccountListAccountidsRequest) GetApiMethodName() string {
    return "alibaba.mozi.buc.account.list.accountids"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMoziBucAccountListAccountidsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ListAccountIds Setter
// 请求参数
func (r *AlibabaMoziBucAccountListAccountidsRequest) SetListAccountIds(_listAccountIds *ListAccountsByAccountIdsRequest) error {
    r._listAccountIds = _listAccountIds
    r.Set("list_account_ids", _listAccountIds)
    return nil
}

// ListAccountIds Getter
func (r AlibabaMoziBucAccountListAccountidsRequest) GetListAccountIds() *ListAccountsByAccountIdsRequest {
    return r._listAccountIds
}
