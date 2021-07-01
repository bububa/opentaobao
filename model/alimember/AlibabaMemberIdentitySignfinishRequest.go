package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
签约确认 API请求
alibaba.member.identity.signfinish

签约确认
*/
type AlibabaMemberIdentitySignfinishAPIRequest struct {
    model.Params
    // 签约确认信息
    _signFinish   *SignIdentityFinishRequest
}

// 初始化AlibabaMemberIdentitySignfinishAPIRequest对象
func NewAlibabaMemberIdentitySignfinishRequest() *AlibabaMemberIdentitySignfinishAPIRequest{
    return &AlibabaMemberIdentitySignfinishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetApiMethodName() string {
    return "alibaba.member.identity.signfinish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SignFinish Setter
// 签约确认信息
func (r *AlibabaMemberIdentitySignfinishAPIRequest) SetSignFinish(_signFinish *SignIdentityFinishRequest) error {
    r._signFinish = _signFinish
    r.Set("sign_finish", _signFinish)
    return nil
}

// SignFinish Getter
func (r AlibabaMemberIdentitySignfinishAPIRequest) GetSignFinish() *SignIdentityFinishRequest {
    return r._signFinish
}
