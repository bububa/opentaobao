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
type TaobaoSimbaRtrptCampaignGetAPIRequest struct {
    model.Params
    // 用户名
    _nick   string
    // 日期，格式yyyy-mm-dd
    _theDate   string
}

// 初始化TaobaoSimbaRtrptCampaignGetAPIRequest对象
func NewTaobaoSimbaRtrptCampaignGetRequest() *TaobaoSimbaRtrptCampaignGetAPIRequest{
    return &TaobaoSimbaRtrptCampaignGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.campaign.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptCampaignGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetNick() string {
    return r._nick
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptCampaignGetAPIRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptCampaignGetAPIRequest) GetTheDate() string {
    return r._theDate
}
