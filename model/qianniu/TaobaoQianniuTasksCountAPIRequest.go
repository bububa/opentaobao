package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTasksCountAPIRequest
任务查询条数接口 API请求
taobao.qianniu.tasks.count

任务查询条数接口 */
type TaobaoQianniuTasksCountAPIRequest struct {
	model.Params
	// 业务类型
	_bizType string
	// 子任务类型
	_subBizType string
	// 业务ID列表，逗号分隔
	_bizIds string
	// 任务的ID列表，用逗号分隔
	_taskIds string
	// 任务发起者用户数字ID
	_senderUid int64
	// 任务执行者用户数字ID
	_receiverUid int64
	// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
	_status string
	// 逗号分隔的子任务状态，由业务方自定义
	_subStatus string
	// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
	_remindFlag int64
	// 任务元id，多个以逗号分隔
	_metadataIds string
	// 与业务相关的买家nick
	_bizNick string
	// 按时间段搜索时的开始日期，格式如2014-01-01，不填则不限
	_startDate string
	// 按时间段搜索的结束日期。不填则不限。格式为2014-01-01
	_endDate string
	// 优先级
	_priority int64
	// 需要排除的任务类型
	_excludeBizType string
	// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
	_keyWord string
}

// New
