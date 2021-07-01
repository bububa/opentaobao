package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRtrptCreativeGetAPIRequest
获取创意实时报表数据 API请求
taobao.simba.rtrpt.creative.get

获取创意实时报表数据 */
type TaobaoSimbaRtrptCreativeGetAPIRequest struct {
	model.Params
	// 用户名
	_nick string
	// 推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 日期，格式yyyy-mm-dd
	_theDate string
}

// New
