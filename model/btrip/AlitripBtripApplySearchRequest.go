package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索审批单 APIRequest
alitrip.btrip.apply.search

外部企业调用获取本企业审批单列表数据
*/
type AlitripBtripApplySearchRequest struct {
    model.Params

    // 请求对象
    rq   *OpenSearchRq 

}

func NewAlitripBtripApplySearchRequest() *AlitripBtripApplySearchRequest{
    return &AlitripBtripApplySearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripApplySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.apply.search"
}

func (r AlitripBtripApplySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripApplySearchRequest) SetRq(rq *OpenSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripApplySearchRequest) GetRq() *OpenSearchRq {
    return r.rq
}

