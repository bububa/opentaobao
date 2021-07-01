package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignScheduleUpdateAPIRequest
更新一个推广计划的分时折扣设置 API请求
taobao.simba.campaign.schedule.update

更新一个推广计划的分时折扣设置 */
type TaobaoSimbaCampaignScheduleUpdateAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 值为：“all”；或者用“;”分割的每天的设置字符串，该字符串为用“,”分割的时段折扣字符串，格式为：起始时间-结束时间:折扣，其中时间是24小时格式记录，折扣是1-150整数，表示折扣百分比；
	_schedule string
	// 主人昵称
	_nick string
}

// New
