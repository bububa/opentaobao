package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskUpdateAPIRequest
更新轻任务 API请求
taobao.qianniu.task.update

由任务执行者调用，sub_status，tag和memo至少提供一个 */
type TaobaoQianniuTaskUpdateAPIRequest struct {
	model.Params
	// 任务ID
	_taskId int64
	// 子任务状态，由业务方自定义
	_subStatus string
	// 任务标签
	_tag string
	// 任务备注。当memo_mode为1时，memo将采用追加方式。
	_memo string
	// 状态值，多个以逗号分隔
	_status string
	// 提醒时间，时间的毫秒数
	_remindTime int64
	// 应用自定义参数
	_bizParam string
	// 0为不提醒，1为全部提醒，2为PC提醒，3为移动提醒，4为已提醒，5为已忽略。
	_remindFlag int64
	// 表示memo字段的更新策略。如需采用追加方式的，请将此字段设置为1。
	_memoMode int64
	// 默认填0，数字越大优化级越高。当前常用0和1.
	_priority int64
	// 0表示没有删除，1表示删除
	_isDeleted int64
}

// New
