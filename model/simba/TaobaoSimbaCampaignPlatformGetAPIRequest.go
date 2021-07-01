package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignPlatformGetAPIRequest
取得一个推广计划的投放平台设置 API请求
taobao.simba.campaign.platform.get

获得一个推广计划的投放平台设置 */
type TaobaoSimbaCampaignPlatformGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// New
