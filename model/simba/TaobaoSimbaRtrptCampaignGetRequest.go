package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广计划实时报表数据 API请求
taobao.simba.rtrpt.campaign.get

获取推广计划实时报表数据
*/
type TaobaoSimbaRtrptCampaignGetRequest struct {
    model.Params
    // 用户名
    nick   string
    // 日期，格式yyyy-mm-dd
    theDate   string
}

// 初始化TaobaoSimbaRtrptCampaignGetRequest对象
func NewTaobaoSimbaRtrptCampaignGetRequest() *TaobaoSimbaRtrptCampaignGetRequest{
    return &TaobaoSimbaRtrptCampaignGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCampaignGetRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCampaignGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptCampaignGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptCampaignGetRequest) GetNick() string {
    return r.nick
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCampaignGetRequest) SetTheDate(theDate string) error {
    r.theDate = theDate
    r.Set("the_date", theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptCampaignGetRequest) GetTheDate() string {
    return r.theDate
}
