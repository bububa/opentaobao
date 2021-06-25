package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户开通的topic列表 APIRequest
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用
*/
type TaobaoTmcUserTopicsGetRequest struct {
    model.Params

    // 卖家nick
    nick   string 

}

func NewTaobaoTmcUserTopicsGetRequest() *TaobaoTmcUserTopicsGetRequest{
    return &TaobaoTmcUserTopicsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTmcUserTopicsGetRequest) GetApiMethodName() string {
    return "taobao.tmc.user.topics.get"
}

func (r TaobaoTmcUserTopicsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTmcUserTopicsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoTmcUserTopicsGetRequest) GetNick() string {
    return r.nick
}

