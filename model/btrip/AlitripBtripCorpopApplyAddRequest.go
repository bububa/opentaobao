package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】isv添加审批单 API请求
alitrip.btrip.corpop.apply.add

【商旅】isv添加审批单
*/
type AlitripBtripCorpopApplyAddRequest struct {
    model.Params
    // 请求参数
    _rq   *OpenApiApplyRq
}

// 初始化AlitripBtripCorpopApplyAddRequest对象
func NewAlitripBtripCorpopApplyAddRequest() *AlitripBtripCorpopApplyAddRequest{
    return &AlitripBtripCorpopApplyAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyAddRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.add"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求参数
func (r *AlitripBtripCorpopApplyAddRequest) SetRq(_rq *OpenApiApplyRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopApplyAddRequest) GetRq() *OpenApiApplyRq {
    return r._rq
}
