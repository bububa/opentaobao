package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加项目 APIRequest
alitrip.btrip.project.add

添加项目
*/
type AlitripBtripProjectAddRequest struct {
    model.Params

    // 请求对象
    rq   *OpenProjectRq 

}

func NewAlitripBtripProjectAddRequest() *AlitripBtripProjectAddRequest{
    return &AlitripBtripProjectAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripProjectAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.project.add"
}

func (r AlitripBtripProjectAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripProjectAddRequest) SetRq(rq *OpenProjectRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripProjectAddRequest) GetRq() *OpenProjectRq {
    return r.rq
}

