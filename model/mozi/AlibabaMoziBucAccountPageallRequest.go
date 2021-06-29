package mozi

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询租户内内所有账号 APIRequest
alibaba.mozi.buc.account.pageall

查询租户内内所有账号
*/
type AlibabaMoziBucAccountPageallRequest struct {
    model.Params

    // 查询租户内所有人员和账号
    pageAll   *PageAllAccountsRequest 

}

func NewAlibabaMoziBucAccountPageallRequest() *AlibabaMoziBucAccountPageallRequest{
    return &AlibabaMoziBucAccountPageallRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMoziBucAccountPageallRequest) GetApiMethodName() string {
    return "alibaba.mozi.buc.account.pageall"
}

func (r AlibabaMoziBucAccountPageallRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMoziBucAccountPageallRequest) SetPageAll(pageAll *PageAllAccountsRequest) error {
    r.pageAll = pageAll
    r.Set("page_all", pageAll)
    return nil
}

func (r AlibabaMoziBucAccountPageallRequest) GetPageAll() *PageAllAccountsRequest {
    return r.pageAll
}

