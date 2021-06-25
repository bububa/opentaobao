package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定会员 APIRequest
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

func NewAlibabaMjMemberBindmemberRequest() *AlibabaMjMemberBindmemberRequest{
    return &AlibabaMjMemberBindmemberRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMjMemberBindmemberRequest) GetApiMethodName() string {
    return "alibaba.mj.member.bindmember"
}

func (r AlibabaMjMemberBindmemberRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMjMemberBindmemberRequest) SetUserId(userId int64) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaMjMemberBindmemberRequest) GetUserId() int64 {
    return r.userId
}

func (r *AlibabaMjMemberBindmemberRequest) SetMallId(mallId int64) error {
    r.mallId = mallId
    r.Set("mall_id", mallId)
    return nil
}

func (r AlibabaMjMemberBindmemberRequest) GetMallId() int64 {
    return r.mallId
}

func (r *AlibabaMjMemberBindmemberRequest) SetOpenId(openId string) error {
    r.openId = openId
    r.Set("open_id", openId)
    return nil
}

func (r AlibabaMjMemberBindmemberRequest) GetOpenId() string {
    return r.openId
}

func (r *AlibabaMjMemberBindmemberRequest) SetChannel(channel string) error {
    r.channel = channel
    r.Set("channel", channel)
    return nil
}

func (r AlibabaMjMemberBindmemberRequest) GetChannel() string {
    return r.channel
}

