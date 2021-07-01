package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveActivityUpdateAPIRequest
互动任务活动修改接口 API请求
taobao.jst.interactive.activity.update

互动任务活动修改接口 */
type TaobaoJstInteractiveActivityUpdateAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
	// 活动开始时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_startTime string
	// 活动结束时间，格式为yyyy-MM-dd HH:mm:ss，任务列表只在活动期间内返回
	_endTime string
	// 活动状态，0=无效，1=有效，status设为0即代表删除此活动，需创建新的活动
	_status int64
}

// New
