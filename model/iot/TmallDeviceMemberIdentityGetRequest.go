package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件会员判断 API请求
tmall.device.member.identity.get

用来识别该用户是否是商家会员·
*/
type TmallDeviceMemberIdentityGetRequest struct {
    model.Params
    // 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
    extraInfo   string
    // 混淆昵称，
    mixNick   string
    // 明文nick，可不填，直接填混淆昵称
    nick   string
}

// 初始化TmallDeviceMemberIdentityGetRequest对象
func NewTmallDeviceMemberIdentityGetRequest() *TmallDeviceMemberIdentityGetRequest{
    return &TmallDeviceMemberIdentityGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceMemberIdentityGetRequest) GetApiMethodName() string {
    return "tmall.device.member.identity.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceMemberIdentityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ExtraInfo Setter
// 扩展参数为JSON字符串，用于埋点统计，source为来源字段固定值 paiyangji代表来源派样机类型设备，deviceId 为设备id，itemId 相关商品id
func (r *TmallDeviceMemberIdentityGetRequest) SetExtraInfo(extraInfo string) error {
    r.extraInfo = extraInfo
    r.Set("extra_info", extraInfo)
    return nil
}

// ExtraInfo Getter
func (r TmallDeviceMemberIdentityGetRequest) GetExtraInfo() string {
    return r.extraInfo
}
// MixNick Setter
// 混淆昵称，
func (r *TmallDeviceMemberIdentityGetRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

// MixNick Getter
func (r TmallDeviceMemberIdentityGetRequest) GetMixNick() string {
    return r.mixNick
}
// Nick Setter
// 明文nick，可不填，直接填混淆昵称
func (r *TmallDeviceMemberIdentityGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TmallDeviceMemberIdentityGetRequest) GetNick() string {
    return r.nick
}
