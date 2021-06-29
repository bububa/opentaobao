package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消确认 APIRequest
alibaba.member.identity.rescindfinish

取消确认
*/
type AlibabaMemberIdentityRescindfinishRequest struct {
    model.Params

    // 取消确认信息
    rescindFinish   *RescindIdentityFinishRequest 

}

func NewAlibabaMemberIdentityRescindfinishRequest() *AlibabaMemberIdentityRescindfinishRequest{
    return &AlibabaMemberIdentityRescindfinishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMemberIdentityRescindfinishRequest) GetApiMethodName() string {
    return "alibaba.member.identity.rescindfinish"
}

func (r AlibabaMemberIdentityRescindfinishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMemberIdentityRescindfinishRequest) SetRescindFinish(rescindFinish *RescindIdentityFinishRequest) error {
    r.rescindFinish = rescindFinish
    r.Set("rescind_finish", rescindFinish)
    return nil
}

func (r AlibabaMemberIdentityRescindfinishRequest) GetRescindFinish() *RescindIdentityFinishRequest {
    return r.rescindFinish
}

