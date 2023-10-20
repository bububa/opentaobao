package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbacampaignbudgetupdateAPIRequest 更新一个推广计划的日限额 API请求
// taobao.simba.campaign.budget.update
//
// 更新一个推广计划的日限额
type TaobaosimbacampaignbudgetupdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
	// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
	_budget int64
	// 是否平滑消耗：false-否，true-是
	_useSmooth bool
}

// NewTaobaosimbacampaignbudgetupdateRequest 初始化TaobaosimbacampaignbudgetupdateAPIRequest对象
func NewTaobaosimbacampaignbudgetupdateRequest() *TaobaosimbacampaignbudgetupdateAPIRequest {
	return &TaobaosimbacampaignbudgetupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.budget.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbacampaignbudgetupdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbacampaignbudgetupdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetBudget is Budget Setter
// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
func (r *TaobaosimbacampaignbudgetupdateAPIRequest) SetBudget(_budget int64) error {
	r._budget = _budget
	r.Set("budget", _budget)
	return nil
}

// GetBudget Budget Getter
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetBudget() int64 {
	return r._budget
}

// SetUseSmooth is UseSmooth Setter
// 是否平滑消耗：false-否，true-是
func (r *TaobaosimbacampaignbudgetupdateAPIRequest) SetUseSmooth(_useSmooth bool) error {
	r._useSmooth = _useSmooth
	r.Set("use_smooth", _useSmooth)
	return nil
}

// GetUseSmooth UseSmooth Getter
func (r TaobaosimbacampaignbudgetupdateAPIRequest) GetUseSmooth() bool {
	return r._useSmooth
}
