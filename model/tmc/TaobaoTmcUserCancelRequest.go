package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取消用户的消息服务 APIRequest
taobao.tmc.user.cancel

取消用户的消息服务
*/
type TaobaoTmcUserCancelRequest struct {
    model.Params

    // 用户昵称
    nick   string 

    // 用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户
    userPlatform   string 

}

func NewTaobaoTmcUserCancelRequest() *TaobaoTmcUserCancelRequest{
    return &TaobaoTmcUserCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcUserCancelRequest) GetApiMethodName() string {
    return "taobao.tmc.user.cancel"
}

func (r TaobaoTmcUserCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcUserCancelRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoTmcUserCancelRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoTmcUserCancelRequest) SetUserPlatform(userPlatform string) error {
    r.userPlatform = userPlatform
    r.Set("user_platform", userPlatform)
    return nil
}

func (r TaobaoTmcUserCancelRequest) GetUserPlatform() string {
    return r.userPlatform
}

