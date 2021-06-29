package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵街会员是否绑定 API请求
alibaba.mj.member.hasbind

喵街检测用户是否为数字化会员
*/
type AlibabaMjMemberHasbindRequest struct {
    model.Params
    // user_id
    userId   int64
    // open_id
    openId   string
}

// 初始化AlibabaMjMemberHasbindRequest对象
func NewAlibabaMjMemberHasbindRequest() *AlibabaMjMemberHasbindRequest{
    return &AlibabaMjMemberHasbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMemberHasbindRequest) GetApiMethodName() string {
    return "alibaba.mj.member.hasbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMemberHasbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// user_id
func (r *AlibabaMjMemberHasbindRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaMjMemberHasbindRequest) GetUserId() int64 {
    return r.userId
}
// OpenId Setter
// open_id
func (r *AlibabaMjMemberHasbindRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

// OpenId Getter
func (r AlibabaMjMemberHasbindRequest) GetOpenId() string {
    return r.openId
}
