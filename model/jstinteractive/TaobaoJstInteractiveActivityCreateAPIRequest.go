package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveActivityCreateAPIRequest
互动任务活动创建接口 API请求
taobao.jst.interactive.activity.create

调用活动创建接口为小程序创建互动任务活动，任务列表仅在活动期间内返回 */
type TaobaoJstInteractiveActivityCreateAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
	// 活动名称
	_activityName string
	// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_startTime string
	// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_endTime string
}

// New
