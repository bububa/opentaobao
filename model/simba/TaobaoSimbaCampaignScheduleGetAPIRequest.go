package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignScheduleGetAPIRequest
取得一个推广计划的分时折扣设置 API请求
taobao.simba.campaign.schedule.get

取得一个推广计划的分时折扣设置 */
type TaobaoSimbaCampaignScheduleGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划Id
	_campaignId int64
}

// New
