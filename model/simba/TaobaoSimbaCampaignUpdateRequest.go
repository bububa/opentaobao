package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划 API请求
taobao.simba.campaign.update

更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
*/
type TaobaoSimbaCampaignUpdateRequest struct {
    model.Params
    // 用户设置的上下限状态；offline-下线；online-上线；
    onlineStatus   string
    // 推广计划Id
    campaignId   int64
    // 推广计划名称，不能多余40个字符，不能和客户其他推广计划同名。
    title   string
    // 主人昵称
    nick   string
}

// 初始化TaobaoSimbaCampaignUpdateRequest对象
func NewTaobaoSimbaCampaignUpdateRequest() *TaobaoSimbaCampaignUpdateRequest{
    return &TaobaoSimbaCampaignUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.campaign.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OnlineStatus Setter
// 用户设置的上下限状态；offline-下线；online-上线；
func (r *TaobaoSimbaCampaignUpdateRequest) SetOnlineStatus(onlineStatus string) error {
    r.onlineStatus = onlineStatus
    r.Set("online_status", onlineStatus)
    return nil
}

// OnlineStatus Getter
func (r TaobaoSimbaCampaignUpdateRequest) GetOnlineStatus() string {
    return r.onlineStatus
}
// CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignUpdateRequest) SetCampaignId(campaignId int64) error {
    r.campaignId = campaignId
    r.Set("campaign_id", campaignId)
    return nil
}

// CampaignId Getter
func (r TaobaoSimbaCampaignUpdateRequest) GetCampaignId() int64 {
    return r.campaignId
}
// Title Setter
// 推广计划名称，不能多余40个字符，不能和客户其他推广计划同名。
func (r *TaobaoSimbaCampaignUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoSimbaCampaignUpdateRequest) GetTitle() string {
    return r.title
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaCampaignUpdateRequest) GetNick() string {
    return r.nick
}
