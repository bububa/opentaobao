package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】查询审批单 APIRequest
alitrip.btrip.corpop.apply.get

【商旅】查询审批单
*/
type AlitripBtripCorpopApplyGetRequest struct {
    model.Params

    // 请求对象
    rq   *OpenIsvSearchRq 

}

func NewAlitripBtripCorpopApplyGetRequest() *AlitripBtripCorpopApplyGetRequest{
    return &AlitripBtripCorpopApplyGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCorpopApplyGetRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.get"
}

func (r AlitripBtripCorpopApplyGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCorpopApplyGetRequest) SetRq(rq *OpenIsvSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCorpopApplyGetRequest) GetRq() *OpenIsvSearchRq {
    return r.rq
}

