package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】搜索审批单列表 APIRequest
alitrip.btrip.corpop.apply.search

【商旅】搜索审批单列表
*/
type AlitripBtripCorpopApplySearchRequest struct {
    model.Params

    // 请求对象
    rq   *OpenIsvSearchRq 

}

func NewAlitripBtripCorpopApplySearchRequest() *AlitripBtripCorpopApplySearchRequest{
    return &AlitripBtripCorpopApplySearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCorpopApplySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.search"
}

func (r AlitripBtripCorpopApplySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCorpopApplySearchRequest) SetRq(rq *OpenIsvSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCorpopApplySearchRequest) GetRq() *OpenIsvSearchRq {
    return r.rq
}

