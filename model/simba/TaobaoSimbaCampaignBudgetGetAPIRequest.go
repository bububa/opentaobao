package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignBudgetGetAPIRequest 取得一个推广计划的日限额 API请求
// taobao.simba.campaign.budget.get
//
// 取得一个推广计划的日限额
type TaobaoSimbaCampaignBudgetGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// NewTaobaoSimbaCampaignBudgetGetRequest 初始化TaobaoSimbaCampaignBudgetGetAPIRequest对象
func NewTaobaoSimbaCampaignBudgetGetRequest() *TaobaoSimbaCampaignBudgetGetAPIRequest {
	return &TaobaoSimbaCampaignBudgetGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaCampaignBudgetGetAPIRequest) Reset() {
	r._nick = ""
	r._campaignId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.campaign.budget.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCampaignBudgetGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetNick() string {
	return r._nick
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaCampaignBudgetGetAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaCampaignBudgetGetAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

var poolTaobaoSimbaCampaignBudgetGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaCampaignBudgetGetRequest()
	},
}

// GetTaobaoSimbaCampaignBudgetGetRequest 从 sync.Pool 获取 TaobaoSimbaCampaignBudgetGetAPIRequest
func GetTaobaoSimbaCampaignBudgetGetAPIRequest() *TaobaoSimbaCampaignBudgetGetAPIRequest {
	return poolTaobaoSimbaCampaignBudgetGetAPIRequest.Get().(*TaobaoSimbaCampaignBudgetGetAPIRequest)
}

// ReleaseTaobaoSimbaCampaignBudgetGetAPIRequest 将 TaobaoSimbaCampaignBudgetGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaCampaignBudgetGetAPIRequest(v *TaobaoSimbaCampaignBudgetGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaCampaignBudgetGetAPIRequest.Put(v)
}
