package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRtrptCampaignGetAPIRequest
获取推广计划实时报表数据 API请求
taobao.simba.rtrpt.campaign.get

获取推广计划实时报表数据 */
type TaobaoSimbaRtrptCampaignGetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// New
