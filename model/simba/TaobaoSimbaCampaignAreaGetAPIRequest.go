package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignAreaGetAPIRequest
取得一个推广计划的投放地域设置 API请求
taobao.simba.campaign.area.get

取得一个推广计划的投放地域设置 */
type TaobaoSimbaCampaignAreaGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// New
