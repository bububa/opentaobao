package btrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除项目 API请求
alitrip.btrip.project.delete

删除项目
*/
type AlitripBtripProjectDeleteRequest struct {
    model.Params
    // 请求对象
    _rq   *OpenProjectRq
}

// 初始化AlitripBtripProjectDeleteRequest对象
func NewAlitripBtripProjectDeleteRequest() *AlitripBtripProjectDeleteRequest{
    return &AlitripBtripProjectDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripBtripProjectDeleteRequest) GetApiMethodName() string {
    return "alitrip.btrip.project.delete"
}

// IRequest interface 方法, 获取API参数
func (r AlitripBtripProjectDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Rq Setter
// 请求对象
func (r *AlitripBtripProjectDeleteRequest) SetRq(_rq *OpenProjectRq) error {
    r._rq = _rq
    r.Set("rq", _rq)
    return nil
}

// Rq Getter
func (r AlitripBtripProjectDeleteRequest) GetRq() *OpenProjectRq {
    return r._rq
}
