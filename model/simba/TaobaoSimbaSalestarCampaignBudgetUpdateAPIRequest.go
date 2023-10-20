package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest 销量明星跟新预算相关接口 API请求
// taobao.simba.salestar.campaign.budget.update
//
// 更新一个推广计划的日限额
type TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
	_budget int64
}

// NewTaobaoSimbaSalestarCampaignBudgetUpdateRequest 初始化TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest对象
func NewTaobaoSimbaSalestarCampaignBudgetUpdateRequest() *TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest {
	return &TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.campaign.budget.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetBudget is Budget Setter
// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
func (r *TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) SetBudget(_budget int64) error {
	r._budget = _budget
	r.Set("budget", _budget)
	return nil
}

// GetBudget Budget Getter
func (r TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest) GetBudget() int64 {
	return r._budget
}
