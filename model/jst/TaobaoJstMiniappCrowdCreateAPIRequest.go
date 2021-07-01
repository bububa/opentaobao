package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstMiniappCrowdCreateAPIRequest
小程序活动创建 API请求
taobao.jst.miniapp.crowd.create

小程序活动创建 */
type TaobaoJstMiniappCrowdCreateAPIRequest struct {
	model.Params
	// 活动开始时间，开始时间和结束时间不能超过1个月
	_endDate string
	// 活动描述
	_description string
	// 活动开始时间
	_startDate string
	// 活动名称
	_crowdName string
}

// New
