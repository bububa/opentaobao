package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignBudgetGetAPIRequest
取得一个推广计划的日限额 API请求
taobao.simba.campaign.budget.get

取得一个推广计划的日限额 */
type TaobaoSimbaCampaignBudgetGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// New
