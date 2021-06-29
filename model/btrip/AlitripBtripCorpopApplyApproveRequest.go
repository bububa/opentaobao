package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】更新审批单状态 API请求
alitrip.btrip.corpop.apply.approve

【商旅】更新审批单状态
*/
type AlitripBtripCorpopApplyApproveRequest struct {
    model.Params
    // 请求对象
    rq   *OpenApiUpdateApplyRq
}

// 初始化AlitripBtripCorpopApplyApproveRequest对象
func NewAlitripBtripCorpopApplyApproveRequest() *AlitripBtripCorpopApplyApproveRequest{
    return &AlitripBtripCorpopApplyApproveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyApproveRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.approve"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyApproveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyApproveRequest) SetRq(rq *OpenApiUpdateApplyRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopApplyApproveRequest) GetRq() *OpenApiUpdateApplyRq {
    return r.rq
}
