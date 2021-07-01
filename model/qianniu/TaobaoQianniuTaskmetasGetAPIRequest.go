package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskmetasGetAPIRequest
任务元查询接口 API请求
taobao.qianniu.taskmetas.get

任务元查询接口 */
type TaobaoQianniuTaskmetasGetAPIRequest struct {
	model.Params
	// 发起任务人的uid
	_senderUid int64
	// 逗号分隔的字段列表.如id,title,content,sender_uid,sender_nick,finish_strategy,biz_sys_Id,biz_sys_task_type,start_time,end_time,reminder_flag,priority
	_fields string
	// 分页数，最大100
	_pageSize int64
	// 当前页码
	_currentPage int64
	// 排序字段。gmt_create,priority等
	_orderBy string
	// 升降序。asc为升，desc为降
	_orderType string
	// 0为未完成。2为完成。4为取消。不填为所有
	_status int64
	// 任务类型
	_bizType string
	// 按关键字搜索
	_keyWord string
	// 客户端的版本信息
	_clientInfo string
	// 接收人uid
	_receiverUid int64
	// 任务元ID，多个以逗号分离
	_metaIds string
}

// New
