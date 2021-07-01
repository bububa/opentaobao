package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaCampaignAddAPIRequest
创建一个推广计划 API请求
taobao.simba.campaign.add

创建一个推广计划 */
type TaobaoSimbaCampaignAddAPIRequest struct {
	model.Params
	// 推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
	_title string
	// 主人昵称
	_nick string
	// 计划类型，当前仅支持两种标准推广0，销量明星16，默认为0
	_type int64
}

// New
