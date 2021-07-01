package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignUpdateAPIRequest
更新一个推广计划 API请求
taobao.simba.campaign.update

更新一个推广计划，可以设置推广计划名字，修改推广计划上下线状态。 */
type TaobaoSimbaCampaignUpdateAPIRequest struct {
	model.Params
	// 用户设置的上下限状态；offline-下线；online-上线；
	_onlineStatus string
	// 推广计划Id
	_campaignId int64
	// 推广计划名称，不能多余40个字符，不能和客户其他推广计划同名。
	_title string
	// 主人昵称
	_nick string
}

// New
