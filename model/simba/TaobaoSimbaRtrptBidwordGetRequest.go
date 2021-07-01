package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广词实时报表数据 API请求
taobao.simba.rtrpt.bidword.get

获取推广词报表数据
*/
type TaobaoSimbaRtrptBidwordGetAPIRequest struct {
    model.Params
    // 用户名
    _nick   string
    // 推广计划id
    _campaignId   int64
    // 推广组id
    _adgroupId   int64
    // 日期，格式yyyy-mm-dd
    _theDate   string
}

// 初始化TaobaoSimbaRtrptBidwordGetAPIRequest对象
func NewTaobaoSimbaRtrptBidwordGetRequest() *TaobaoSimbaRtrptBidwordGetAPIRequest{
    return &TaobaoSimbaRtrptBidwordGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetApiMethodName() string {
    return "taobao.simba.rtrpt.bidword.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 用户名
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetNick() string {
    return r._nick
}
// CampaignId Setter
// 推广计划id
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetCampaignId(_campaignId int64) error {
    r._campaignId = _campaignId
    r.Set("campaign_id", _campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetCampaignId() int64 {
    return r._campaignId
}
// AdgroupId Setter
// 推广组id
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
// TheDate Setter
// 日期，格式yyyy-mm-dd
func (r *TaobaoSimbaRtrptBidwordGetAPIRequest) SetTheDate(_theDate string) error {
    r._theDate = _theDate
    r.Set("the_date", _theDate)
    return nil
}

// TheDate Getter
func (r TaobaoSimbaRtrptBidwordGetAPIRequest) GetTheDate() string {
    return r._theDate
}
