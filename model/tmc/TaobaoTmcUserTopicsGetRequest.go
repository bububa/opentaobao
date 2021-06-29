package tmc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户开通的topic列表 API请求
taobao.tmc.user.topics.get

获取用户开通的topic列表，授权消息使用
*/
type TaobaoTmcUserTopicsGetRequest struct {
    model.Params
    // 卖家nick
    nick   string
}

// 初始化TaobaoTmcUserTopicsGetRequest对象
func NewTaobaoTmcUserTopicsGetRequest() *TaobaoTmcUserTopicsGetRequest{
    return &TaobaoTmcUserTopicsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserTopicsGetRequest) GetApiMethodName() string {
    return "taobao.tmc.user.topics.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserTopicsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 卖家nick
func (r *TaobaoTmcUserTopicsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoTmcUserTopicsGetRequest) GetNick() string {
    return r.nick
}
