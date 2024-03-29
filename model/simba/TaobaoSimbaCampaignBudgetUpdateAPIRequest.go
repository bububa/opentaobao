package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignBudgetUpdateAPIRequest 更新一个推广计划的日限额 API请求
// taobao.simba.campaign.budget.update
//
// 更新一个推广计划的日限额
type TaobaoSimbaCampaignBudgetUpdateAPIRequest struct {
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

// NewTaobaoSimbaCampaignBudgetUpdateRequest 初始化TaobaoSimbaCampaignBudgetUpdateAPIRequest对象
func NewTaobaoSimbaCampaignBudgetUpdateRequest() *TaobaoSimbaCampaignBudgetUpdateAPIRequest {
	return &TaobaoSimbaCampaignBudgetUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignBudgetUpdateAPIRequest) Reset() {
	r._nick = ""
	r._campaignId = 0
	r._budget = 0
	r._useSmooth = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.budget.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignBudgetUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignBudgetUpdateAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetBudget is Budget Setter
// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
func (r *TaobaoSimbaCampaignBudgetUpdateAPIRequest) SetBudget(_budget int64) error {
	r._budget = _budget
	r.Set("budget", _budget)
	return nil
}

// GetBudget Budget Getter
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetBudget() int64 {
	return r._budget
}

// SetUseSmooth is UseSmooth Setter
// 是否平滑消耗：false-否，true-是
func (r *TaobaoSimbaCampaignBudgetUpdateAPIRequest) SetUseSmooth(_useSmooth bool) error {
	r._useSmooth = _useSmooth
	r.Set("use_smooth", _useSmooth)
	return nil
}

// GetUseSmooth UseSmooth Getter
func (r TaobaoSimbaCampaignBudgetUpdateAPIRequest) GetUseSmooth() bool {
	return r._useSmooth
}

var poolTaobaoSimbaCampaignBudgetUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignBudgetUpdateRequest()
	},
}

// GetTaobaoSimbaCampaignBudgetUpdateRequest 从 sync.Pool 获取 TaobaoSimbaCampaignBudgetUpdateAPIRequest
func GetTaobaoSimbaCampaignBudgetUpdateAPIRequest() *TaobaoSimbaCampaignBudgetUpdateAPIRequest {
	return poolTaobaoSimbaCampaignBudgetUpdateAPIRequest.Get().(*TaobaoSimbaCampaignBudgetUpdateAPIRequest)
}

// ReleaseTaobaoSimbaCampaignBudgetUpdateAPIRequest 将 TaobaoSimbaCampaignBudgetUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignBudgetUpdateAPIRequest(v *TaobaoSimbaCampaignBudgetUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignBudgetUpdateAPIRequest.Put(v)
}
