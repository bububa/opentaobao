package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定会员 API请求
alibaba.mj.member.bindmember

用于绑定喵街数字化会员
*/
type AlibabaMjMemberBindmemberRequest struct {
    model.Params
    // 用户号
    userId   int64
    // 商城Id
    mallId   int64
    // open_id
    openId   string
    // 渠道
    channel   string
}

// 初始化AlibabaMjMemberBindmemberRequest对象
func NewAlibabaMjMemberBindmemberRequest() *AlibabaMjMemberBindmemberRequest{
    return &AlibabaMjMemberBindmemberRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMjMemberBindmemberRequest) GetApiMethodName() string {
    return "alibaba.mj.member.bindmember"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMjMemberBindmemberRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户号
func (r *AlibabaMjMemberBindmemberRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaMjMemberBindmemberRequest) GetUserId() int64 {
    return r.userId
}
// MallId Setter
// 商城Id
func (r *AlibabaMjMemberBindmemberRequest) SetMallId(mallId int64) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

// MallId Getter
func (r AlibabaMjMemberBindmemberRequest) GetMallId() int64 {
    return r.mallId
}
// OpenId Setter
// open_id
func (r *AlibabaMjMemberBindmemberRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

// OpenId Getter
func (r AlibabaMjMemberBindmemberRequest) GetOpenId() string {
    return r.openId
}
// Channel Setter
// 渠道
func (r *AlibabaMjMemberBindmemberRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

// Channel Getter
func (r AlibabaMjMemberBindmemberRequest) GetChannel() string {
    return r.channel
}
