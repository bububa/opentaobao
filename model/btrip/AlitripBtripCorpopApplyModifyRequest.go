package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】修改出差审批单（行程） APIRequest
alitrip.btrip.corpop.apply.modify

【商旅】修改出差审批单（行程）
*/
type AlitripBtripCorpopApplyModifyRequest struct {
    model.Params

    // 请求对象
    rq   *OpenApiApplyRq 

}

func NewAlitripBtripCorpopApplyModifyRequest() *AlitripBtripCorpopApplyModifyRequest{
    return &AlitripBtripCorpopApplyModifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCorpopApplyModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.modify"
}

func (r AlitripBtripCorpopApplyModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCorpopApplyModifyRequest) SetRq(rq *OpenApiApplyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCorpopApplyModifyRequest) GetRq() *OpenApiApplyRq {
    return r.rq
}

