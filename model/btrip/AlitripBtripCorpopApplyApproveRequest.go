package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】更新审批单状态 APIRequest
alitrip.btrip.corpop.apply.approve

【商旅】更新审批单状态
*/
type AlitripBtripCorpopApplyApproveRequest struct {
    model.Params

    // 请求对象
    rq   *OpenApiUpdateApplyRq 

}

func NewAlitripBtripCorpopApplyApproveRequest() *AlitripBtripCorpopApplyApproveRequest{
    return &AlitripBtripCorpopApplyApproveRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripBtripCorpopApplyApproveRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.approve"
}

func (r AlitripBtripCorpopApplyApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripBtripCorpopApplyApproveRequest) SetRq(rq *OpenApiUpdateApplyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

func (r AlitripBtripCorpopApplyApproveRequest) GetRq() *OpenApiUpdateApplyRq {
    return r.rq
}

