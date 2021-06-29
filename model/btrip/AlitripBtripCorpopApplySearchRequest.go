package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】搜索审批单列表 API请求
alitrip.btrip.corpop.apply.search

【商旅】搜索审批单列表
*/
type AlitripBtripCorpopApplySearchRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenIsvSearchRq
}

// 初始化AlitripBtripCorpopApplySearchRequest对象
func NewAlitripBtripCorpopApplySearchRequest() *AlitripBtripCorpopApplySearchRequest{
    return &AlitripBtripCorpopApplySearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripCorpopApplySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.corpop.apply.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripCorpopApplySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripCorpopApplySearchRequest) SetRq(_rq *OpenIsvSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripCorpopApplySearchRequest) GetRq() *OpenIsvSearchRq {
    return r._rq
}
