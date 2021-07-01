package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptTargetingtagbaseGetAPIRequest
定向基础报表 API请求
taobao.simba.rpt.targetingtagbase.get

获取定向基础报表 */
type TaobaoSimbaRptTargetingtagbaseGetAPIRequest struct {
	model.Params
	// 被操作者昵称
	_nick string
	// 计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 起始时间
	_startTime string
	// 结束时间
	_endTime string
	// 分页大小
	_pageSize int64
	// 页码
	_pageNumber int64
}

// New
