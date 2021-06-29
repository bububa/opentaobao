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
type AlitripBtripApplySearchRequest struct {
    model.Params
    // 请求对象
    rq   *OpenSearchRq
}

// 初始化AlitripBtripApplySearchRequest对象
func NewAlitripBtripApplySearchRequest() *AlitripBtripApplySearchRequest{
    return &AlitripBtripApplySearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripApplySearchRequest) GetApiMethodName() string {
    return "alitrip.btrip.apply.search"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripApplySearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripApplySearchRequest) SetRq(rq *OpenSearchRq) error {
    r.rq = rq
    r.Set("rq", rq)
    return nil
}

// Rq Getter
func (r AlitripBtripApplySearchRequest) GetRq() *OpenSearchRq {
    return r.rq
}
