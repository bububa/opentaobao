package idleparttime

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleParttimeNotifyAPIRequest
兼职通知接口 API请求
alibaba.idle.parttime.notify

服务商侧有岗位状态变更对我们进行通知(岗位关闭, 录取状态) */
type AlibabaIdleParttimeNotifyAPIRequest struct {
	model.Params
	// 实时同步类型, 0: 岗位状态, 1: 录取状态
	_type int64
	// 岗位: 0关闭 ; 录取: 0不录取, 1录取
	_status int64
	// 岗位id
	_jobId int64
	// 用户id
	_userId int64
	// 同步时间
	_syncTime int64
	// 报名id
	_applyId int64
	// 通知消息
	_message string
}

// New
