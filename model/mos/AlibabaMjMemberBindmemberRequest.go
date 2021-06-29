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
    _userId   int64
    // 商城Id
    _mallId   int64
    // open_id
    _openId   string
    // 渠道
    _channel   string
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
func (r *AlibabaMjMemberBindmemberRequest) SetUserId(_userId int64) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaMjMemberBindmemberRequest) GetUserId() int64 {
    return r._userId
}
// MallId Setter
// 商城Id
func (r *AlibabaMjMemberBindmemberRequest) SetMallId(_mallId int64) error {
    r._mallId = _mallId
    r.Set("mall_id", _mallId)
    return nil
}

// MallId Getter
func (r AlibabaMjMemberBindmemberRequest) GetMallId() int64 {
    return r._mallId
}
// OpenId Setter
// open_id
func (r *AlibabaMjMemberBindmemberRequest) SetOpenId(_openId string) error {
    r._openId = _openId
    r.Set("open_id", _openId)
    return nil
}

// OpenId Getter
func (r AlibabaMjMemberBindmemberRequest) GetOpenId() string {
    return r._openId
}
// Channel Setter
// 渠道
func (r *AlibabaMjMemberBindmemberRequest) SetChannel(_channel string) error {
    r._channel = _channel
    r.Set("channel", _channel)
    return nil
}

// Channel Getter
func (r AlibabaMjMemberBindmemberRequest) GetChannel() string {
    return r._channel
}
