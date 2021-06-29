package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
变更项目 API请求
alitrip.btrip.project.modify

变更项目
*/
type AlitripBtripProjectModifyRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenProjectRq
}

// 初始化AlitripBtripProjectModifyRequest对象
func NewAlitripBtripProjectModifyRequest() *AlitripBtripProjectModifyRequest{
    return &AlitripBtripProjectModifyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectModifyRequest) GetApiMethodName() string {
    return "alitrip.btrip.project.modify"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectModifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripProjectModifyRequest) SetRq(_rq *OpenProjectRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripProjectModifyRequest) GetRq() *OpenProjectRq {
    return r._rq
}
