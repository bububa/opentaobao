package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广计划 APIRequest
taobao.simba.campaign.add

创建一个推广计划
*/
type TaobaoSimbaCampaignAddRequest struct {
    model.Params

    // 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
    title   string 

    // 主人昵称
    nick   string 

    // 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
    type   int64 

}

func NewTaobaoSimbaCampaignAddRequest() *TaobaoSimbaCampaignAddRequest{
    return &TaobaoSimbaCampaignAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCampaignAddRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.add"
}

func (r TaobaoSimbaCampaignAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCampaignAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r TaobaoSimbaCampaignAddRequest) GetTitle() string {
    return r.title
}

func (r *TaobaoSimbaCampaignAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCampaignAddRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCampaignAddRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

func (r TaobaoSimbaCampaignAddRequest) GetType() int64 {
    return r.type
}

