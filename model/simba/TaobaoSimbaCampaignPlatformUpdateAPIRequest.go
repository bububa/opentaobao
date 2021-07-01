package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignPlatformUpdateAPIRequest
更新一个推广计划的平台设置 API请求
taobao.simba.campaign.platform.update

更新一个推广计划的平台设置 */
type TaobaoSimbaCampaignPlatformUpdateAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 搜索投放频道代码数组，频道代码必须是直通车搜索类频道列表中的值。1：淘宝站内搜索，8、无线站内搜索；16:无线站外搜索
	_searchChannels []int64
	// 非搜索投放频道代码数组，频道代码必须是直通车非搜索类频道列表中的值。1、淘宝站内定向；2、站外定向；8、无线站内定向；16、无线站外定向
	_nonsearchChannels []int64
	// 已经废弃
	_outsideDiscount int64
	// 已经废弃
	_mobileDiscount int64
	// 主人昵称
	_nick string
}

// New
