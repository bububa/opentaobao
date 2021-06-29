package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
变更项目 APIRequest
alitrip.btrip.project.modify

变更项目
*/
type AlitripBtripProjectModifyRequest struct {
    model.Params

    // 请求对象
    rq   *OpenProjectRq 

}

func NewAlitripBtripProjectModifyRequest() *AlitripBtripProjectModifyRequest{
    return &AlitripBtripProjectModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripProjectModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.project.modify"
}

func (r AlitripBtripProjectModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripProjectModifyRequest) SetRq(rq *OpenProjectRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripProjectModifyRequest) GetRq() *OpenProjectRq {
    return r.rq
}

