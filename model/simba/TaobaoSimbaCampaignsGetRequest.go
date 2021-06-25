package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一组推广计划 APIRequest
taobao.simba.campaigns.get

取得一个客户的推广计划；
*/
type TaobaoSimbaCampaignsGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 计划类型0位标准计划，16位销量明星计划
    type   int64 

}

func NewTaobaoSimbaCampaignsGetRequest() *TaobaoSimbaCampaignsGetRequest{
    return &TaobaoSimbaCampaignsGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignsGetRequest) GetApiMethodName() string {
    return "taobao.simba.campaigns.get"
}

func (r TaobaoSimbaCampaignsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignsGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignsGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCampaignsGetRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoSimbaCampaignsGetRequest) GetType() int64 {
    return r.type
}

