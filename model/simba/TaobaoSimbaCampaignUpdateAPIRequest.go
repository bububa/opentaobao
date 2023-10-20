package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignupdateAPIRequest 更新一个推广计划 API请求
// taobao.simba.campaign.update
//
// 更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。
type TaobaosimbacampaignupdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划名称，不能多余40个字符，不能和客户其他推广计划同名。
	_title string
	// 用户设置的上下限状态；offline-下线；online-上线；
	_onlineStatus string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaosimbacampaignupdateRequest 初始化TaobaosimbacampaignupdateAPIRequest对象
func NewTaobaosimbacampaignupdateRequest() *TaobaosimbacampaignupdateAPIRequest {
	return &TaobaosimbacampaignupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignupdateAPIRequest) GetNick() string {
	return r._nick
}

// SetTitle is Title Setter
// 推广计划名称，不能多余40个字符，不能和客户其他推广计划同名。
func (r *TaobaosimbacampaignupdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaosimbacampaignupdateAPIRequest) GetTitle() string {
	return r._title
}

// SetOnlineStatus is OnlineStatus Setter
// 用户设置的上下限状态；offline-下线；online-上线；
func (r *TaobaosimbacampaignupdateAPIRequest) SetOnlineStatus(_onlineStatus string) error {
	r._onlineStatus = _onlineStatus
	r.Set("online_status", _onlineStatus)
	return nil
}

// GetOnlineStatus OnlineStatus Getter
func (r TaobaosimbacampaignupdateAPIRequest) GetOnlineStatus() string {
	return r._onlineStatus
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignupdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignupdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}
