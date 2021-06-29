package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】isv添加审批单 APIRequest
alitrip.btrip.corpop.apply.add

【商旅】isv添加审批单
*/
type AlitripBtripCorpopApplyAddRequest struct {
    model.Params

    // 请求参数
    rq   *OpenApiApplyRq 

}

func NewAlitripBtripCorpopApplyAddRequest() *AlitripBtripCorpopApplyAddRequest{
    return &AlitripBtripCorpopApplyAddRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCorpopApplyAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.add"
}

func (r AlitripBtripCorpopApplyAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCorpopApplyAddRequest) SetRq(rq *OpenApiApplyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCorpopApplyAddRequest) GetRq() *OpenApiApplyRq {
    return r.rq
}

