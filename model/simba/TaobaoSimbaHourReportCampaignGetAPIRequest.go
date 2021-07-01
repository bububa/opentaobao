package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaHourReportCampaignGetAPIRequest
计划维度小时报表获取 API请求
taobao.simba.hour.report.campaign.get

计划维度小时报表获取 */
type TaobaoSimbaHourReportCampaignGetAPIRequest struct {
	model.Params
	// 昵称
	_nick string
	// 时间
	_theDate string
	// 当前小时
	_hour string
	// 计划id
	_campaignId string
}

// New
