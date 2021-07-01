package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTasksGetAPIRequest
获取指定的任务 API请求
taobao.qianniu.tasks.get

获取指定的任务，可用的参数组合：<br/>task_ids + need_meta + fields：精确查找<br/>biz_type + sub_biz_type + biz_ids + need_meta + fields：按照业务ID查找<br/>biz_type + sub_biz_type + sender_uid + need_meta + fields：按照发起者查找<br/>biz_type + sub_biz_type + receiver_uid + need_meta + fields：按照执行者查找<br/>biz_type+modify_start_time+modify_end_time+fields:能支持指定修改时间的查询，用于增量查询等 */
type TaobaoQianniuTasksGetAPIRequest struct {
	model.Params
	// 排序字段，可以为id,gmt_create,gmt_finished,metadata_id等
	_orderBy string
	// asc为升，desc为降
	_orderType string
	// 0-不需要提醒，未设提醒时间 1-设置过提醒时间，需要提醒
	_remindFlag int64
	// 业务相关的对象，当前主要表示买家nick
	_bizNick string
	// 根据任务创建时间搜索的开始日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
	_startDate string
	// 根据任务创建时间搜索的结束日期（含），不填则不限。例如只查询2014-01-01当天的任务，则将start_date和end_date都设置成2014-01-01
	_endDate string
	// 根据任务修改时间搜索的开始时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之后有修改的任务，则将modify_start_time_str设置成“2014-01-01 00:00:10”
	_modifyStartTimeStr string
	// 根据任务修改时间搜索的结束时间（含），不填则不限。例如查询“2014-01-01 00:00:10”之前有修改的任务，则将modify_end_time_str设置成“2014-01-01 00:00:10”
	_modifyEndTimeStr string
	// 优先级。即创建时的metadata中的优先级。0为低，1为中，2为高。
	_priority int64
	// 需要排除的任务类型
	_excludeBizType string
	// 关键词搜索。只对任务内容进行模糊匹配，以及bizid和biznick进行精准匹配
	_keyWord string
	// 当前页数，从1开始
	_currentPage int64
	// 每页条数
	_pageSize int64
	// 业务类型
	_bizType string
	// 子任务类型
	_subBizType string
	// 任务的ID列表，用逗号分隔
	_taskIds string
	// 业务ID列表，逗号分隔
	_bizIds string
	// 任务执行者用户数字ID
	_receiverUid int64
	// 任务发起者用户数字ID
	_senderUid int64
	// 逗号分隔的任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略
	_status string
	// 逗号分隔的子任务状态，由业务方自定义
	_subStatus string
	// 任务元id，多个以逗号分隔
	_metadataIds string
	// 是否需要meta信息，默认值为false
	_needMeta bool
	// 逗号分隔的字段列表，各个字段含义： id：任务ID receiver_uid：执行者用户数字ID receiver_nick：执行者用户昵称 status：任务状态：0-未执行，1-执行中，2-执行完成，3-超时，4-取消，5-忽略 sub_status：子任务状态，由业务方自定义 finish_strategy：任务完成策略：1-一个人完成，2-所有人完成 gmt_finished：任务完成时间，格式：时间毫秒数 biz_type：业务类型 sub_biz_type：子业务类型 biz_id：业务ID biz_param：业务参数 biz_entry：业务入口 tag：任务标签 memo：任务备注
	_fields string
	// 客户端的版本信息
	_clientInfo string
	// 是否需要删除的任务，默认为false
	_needDeleted bool
}

// New
