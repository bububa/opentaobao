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
type AlitripBtripCorpopApplyModifyAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenApiApplyRq
}

// 初始化AlitripBtripCorpopApplyModifyAPIRequest对象
func NewAlitripBtripCorpopApplyModifyRequest() *AlitripBtripCorpopApplyModifyAPIRequest{
    return &AlitripBtripCorpopApplyModifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyModifyAPIRequest) SetRq(_rq *OpenApiApplyRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopApplyModifyAPIRequest) GetRq() *OpenApiApplyRq {
    return r._rq
}
