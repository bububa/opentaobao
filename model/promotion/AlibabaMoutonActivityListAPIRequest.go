package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoutonActivityListAPIRequest
商家自运营活动列表 API请求
alibaba.mouton.activity.list

商家查询自己配置的活动列表 */
type AlibabaMoutonActivityListAPIRequest struct {
	model.Params
	// 开始时间
	_startTimeEnd string
	// 每页记录数
	_pageSize int64
	// 来源
	_source string
	// 开始时间
	_startTimeBegin string
	// 结束时间
	_endTimeBegin string
	// 结束时间
	_endTimeEnd string
	// 来源记录id
	_sourceRecordId int64
	// 状态
	_statusList []string
	// 当前页
	_currentPage int64
}

// New
