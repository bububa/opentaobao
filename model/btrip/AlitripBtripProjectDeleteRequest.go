package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除项目 APIRequest
alitrip.btrip.project.delete

删除项目
*/
type AlitripBtripProjectDeleteRequest struct {
    model.Params

    // 请求对象
    rq   *OpenProjectRq 

}

func NewAlitripBtripProjectDeleteRequest() *AlitripBtripProjectDeleteRequest{
    return &AlitripBtripProjectDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripProjectDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.project.delete"
}

func (r AlitripBtripProjectDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripProjectDeleteRequest) SetRq(rq *OpenProjectRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripProjectDeleteRequest) GetRq() *OpenProjectRq {
    return r.rq
}

