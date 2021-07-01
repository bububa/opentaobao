package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索审批单 API请求
alitrip.btrip.apply.search

外部企业调用获取本企业审批单列表数据
*/
type AlitripBtripApplySearchAPIRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenSearchRq
}

// 初始化AlitripBtripApplySearchAPIRequest对象
func NewAlitripBtripApplySearchRequest() *AlitripBtripApplySearchAPIRequest{
    return &AlitripBtripApplySearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApplySearchAPIRequest) GetApiMethodName() string {
    return "alitrip.btrip.apply.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApplySearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripApplySearchAPIRequest) SetRq(_rq *OpenSearchRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripApplySearchAPIRequest) GetRq() *OpenSearchRq {
    return r._rq
}
