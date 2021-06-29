package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场次接口pushPerform APIRequest
alibaba.damai.mev.open.pushperform

pushPerform
*/
type AlibabaDamaiMevOpenPushperformRequest struct {
    model.Params

    // 入参pushPerformParam
    pushPerformParam   *ThirdPerformPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushperformRequest() *AlibabaDamaiMevOpenPushperformRequest{
    return &AlibabaDamaiMevOpenPushperformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushperformRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushperform"
}

func (r AlibabaDamaiMevOpenPushperformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushperformRequest) SetPushPerformParam(pushPerformParam *ThirdPerformPushOpenParam) error {
    r.pushPerformParam = pushPerformParam
    r.Set("push_perform_param", pushPerformParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushperformRequest) GetPushPerformParam() *ThirdPerformPushOpenParam {
    return r.pushPerformParam
}

