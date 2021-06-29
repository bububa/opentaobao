package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消确认 API请求
alibaba.member.identity.rescindfinish

取消确认
*/
type AlibabaMemberIdentityRescindfinishRequest struct {
    model.Params
    // 取消确认信息
    _rescindFinish   *RescindIdentityFinishRequest
}

// 初始化AlibabaMemberIdentityRescindfinishRequest对象
func NewAlibabaMemberIdentityRescindfinishRequest() *AlibabaMemberIdentityRescindfinishRequest{
    return &AlibabaMemberIdentityRescindfinishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentityRescindfinishRequest) GetApiMethodName() string {
    return "alibaba.member.identity.rescindfinish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentityRescindfinishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RescindFinish Setter
// 取消确认信息
func (r *AlibabaMemberIdentityRescindfinishRequest) SetRescindFinish(_rescindFinish *RescindIdentityFinishRequest) error {
    r._rescindFinish = _rescindFinish
    r.Set("rescind_finish", _rescindFinish)
    return nil
}

// RescindFinish Getter
func (r AlibabaMemberIdentityRescindfinishRequest) GetRescindFinish() *RescindIdentityFinishRequest {
    return r._rescindFinish
}
