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
    _extraInfo   string
    // 混淆昵称，
    _mixNick   string
    // 明文nick，可不填，直接填混淆昵称
    _nick   string
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
func (r *TmallDeviceMemberIdentityGetRequest) SetExtraInfo(_extraInfo string) error {
    r._extraInfo = _extraInfo
    r.Set("extra_info", _extraInfo)
    return nil
}

// ExtraInfo Getter
func (r TmallDeviceMemberIdentityGetRequest) GetExtraInfo() string {
    return r._extraInfo
}
// MixNick Setter
// 混淆昵称，
func (r *TmallDeviceMemberIdentityGetRequest) SetMixNick(_mixNick string) error {
    r._mixNick = _mixNick
    r.Set("mix_nick", _mixNick)
    return nil
}

// MixNick Getter
func (r TmallDeviceMemberIdentityGetRequest) GetMixNick() string {
    return r._mixNick
}
// Nick Setter
// 明文nick，可不填，直接填混淆昵称
func (r *TmallDeviceMemberIdentityGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TmallDeviceMemberIdentityGetRequest) GetNick() string {
    return r._nick
}
