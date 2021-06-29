package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广计划 API请求
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

// 初始化TaobaoSimbaCampaignAddRequest对象
func NewTaobaoSimbaCampaignAddRequest() *TaobaoSimbaCampaignAddRequest{
    return &TaobaoSimbaCampaignAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignAddRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Title Setter
// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
func (r *TaobaoSimbaCampaignAddRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoSimbaCampaignAddRequest) GetTitle() string {
    return r.title
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignAddRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignAddRequest) GetNick() string {
    return r.nick
}
// Type Setter
// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
func (r *TaobaoSimbaCampaignAddRequest) SetType(type int64) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r TaobaoSimbaCampaignAddRequest) GetType() int64 {
    return r.type
}
