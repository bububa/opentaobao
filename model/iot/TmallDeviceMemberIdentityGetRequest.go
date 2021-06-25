package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
智能硬件会员判断 APIRequest
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

func NewTmallDeviceMemberIdentityGetRequest() *TmallDeviceMemberIdentityGetRequest{
    return &TmallDeviceMemberIdentityGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallDeviceMemberIdentityGetRequest) GetApiMethodName() string {
    return "tmall.device.member.identity.get"
}

func (r TmallDeviceMemberIdentityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallDeviceMemberIdentityGetRequest) SetExtraInfo(extraInfo string) error {
    r.extraInfo = extraInfo
    r.Set("extra_info", extraInfo)
    return nil
}

func (r TmallDeviceMemberIdentityGetRequest) GetExtraInfo() string {
    return r.extraInfo
}

func (r *TmallDeviceMemberIdentityGetRequest) SetMixNick(mixNick string) error {
    r.mixNick = mixNick
    r.Set("mix_nick", mixNick)
    return nil
}

func (r TmallDeviceMemberIdentityGetRequest) GetMixNick() string {
    return r.mixNick
}

func (r *TmallDeviceMemberIdentityGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TmallDeviceMemberIdentityGetRequest) GetNick() string {
    return r.nick
}

