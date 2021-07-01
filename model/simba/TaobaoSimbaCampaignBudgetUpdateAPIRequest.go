package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignBudgetUpdateAPIRequest
更新一个推广计划的日限额 API请求
taobao.simba.campaign.budget.update

更新一个推广计划的日限额 */
type TaobaoSimbaCampaignBudgetUpdateAPIRequest struct {
	model.Params
	// 是否平滑消耗：false-否，true-是
	_useSmooth bool
	// 推广计划Id
	_campaignId int64
	// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
	_budget int64
	// 主人昵称
	_nick string
}

// New
