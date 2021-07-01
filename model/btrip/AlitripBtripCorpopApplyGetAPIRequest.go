package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】查询审批单 API请求
alitrip.btrip.corpop.apply.get

【商旅】查询审批单
*/
type AlitripBtripCorpopApplyGetAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenIsvSearchRq
}

// 初始化AlitripBtripCorpopApplyGetAPIRequest对象
func NewAlitripBtripCorpopApplyGetRequest() *AlitripBtripCorpopApplyGetAPIRequest{
    return &AlitripBtripCorpopApplyGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplyGetAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplyGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplyGetAPIRequest) SetRq(_rq *OpenIsvSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopApplyGetAPIRequest) GetRq() *OpenIsvSearchRq {
    return r._rq
}
