package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】修改出差审批单（行程） API请求
alitrip.btrip.corpop.apply.modify

【商旅】修改出差审批单（行程）
*/
type AlitripBtripCorpopApplyModifyRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenApiApplyRq
}

// 初始化AlitripBtripCorpopApplyModifyRequest对象
func NewAlitripBtripCorpopApplyModifyRequest() *AlitripBtripCorpopApplyModifyRequest{
    return &AlitripBtripCorpopApplyModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyModifyRequest) SetRq(_rq *OpenApiApplyRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopApplyModifyRequest) GetRq() *OpenApiApplyRq {
    return r._rq
}
