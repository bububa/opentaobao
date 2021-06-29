package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签约确认 APIRequest
alibaba.member.identity.signfinish

签约确认
*/
type AlibabaMemberIdentitySignfinishRequest struct {
    model.Params

    // 签约确认信息
    signFinish   *SignIdentityFinishRequest 

}

func NewAlibabaMemberIdentitySignfinishRequest() *AlibabaMemberIdentitySignfinishRequest{
    return &AlibabaMemberIdentitySignfinishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMemberIdentitySignfinishRequest) GetApiMethodName() string {
    return "alibaba.member.identity.signfinish"
}

func (r AlibabaMemberIdentitySignfinishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMemberIdentitySignfinishRequest) SetSignFinish(signFinish *SignIdentityFinishRequest) error {
    r.signFinish = signFinish
    r.Set("sign_finish", signFinish)
    return nil
}

func (r AlibabaMemberIdentitySignfinishRequest) GetSignFinish() *SignIdentityFinishRequest {
    return r.signFinish
}

