package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignAreaUpdateAPIRequest
更新一个推广计划的投放地域 API请求
taobao.simba.campaign.area.update

更新一个推广计划的投放地域 */
type TaobaoSimbaCampaignAreaUpdateAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 值为：“all”；或者用“,”分割的数字，数字必须是直通车全国省市列表的AreaID；
	_area string
	// 主人昵称
	_nick string
}

// New
