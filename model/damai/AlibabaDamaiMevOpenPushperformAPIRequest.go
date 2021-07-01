package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场次接口pushPerform API请求
alibaba.damai.mev.open.pushperform

pushPerform
*/
type AlibabaDamaiMevOpenPushperformAPIRequest struct {
    model.Params
    // 入参pushPerformParam
    _pushPerformParam   *ThirdPerformPushOpenParam
}

// 初始化AlibabaDamaiMevOpenPushperformAPIRequest对象
func NewAlibabaDamaiMevOpenPushperformRequest() *AlibabaDamaiMevOpenPushperformAPIRequest{
    return &AlibabaDamaiMevOpenPushperformAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushperform"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushPerformParam Setter
// 入参pushPerformParam
func (r *AlibabaDamaiMevOpenPushperformAPIRequest) SetPushPerformParam(_pushPerformParam *ThirdPerformPushOpenParam) error {
    r._pushPerformParam = _pushPerformParam
    r.Set("push_perform_param", _pushPerformParam)
    return nil
}

// PushPerformParam Getter
func (r AlibabaDamaiMevOpenPushperformAPIRequest) GetPushPerformParam() *ThirdPerformPushOpenParam {
    return r._pushPerformParam
}
