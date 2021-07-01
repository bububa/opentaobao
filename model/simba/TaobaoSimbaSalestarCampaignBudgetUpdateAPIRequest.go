package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest
销量明星跟新预算相关接口 API请求
taobao.simba.salestar.campaign.budget.update

更新一个推广计划的日限额 */
type TaobaoSimbaSalestarCampaignBudgetUpdateAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 如果为空则取消限额；否则必须为整数，单位是元，不得小于30；
	_budget int64
}

// New
