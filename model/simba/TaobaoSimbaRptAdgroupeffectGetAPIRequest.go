package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptAdgroupeffectGetAPIRequest
推广组效果报表数据对象 API请求
taobao.simba.rpt.adgroupeffect.get

推广组效果报表数据对象 */
type TaobaoSimbaRptAdgroupeffectGetAPIRequest struct {
	model.Params
	// 权限校验参数
	_subwayToken string
	// 昵称
	_nick string
	// 推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 报表类型（搜索：SEARCH,类目出价：CAT,<br/>定向投放：NOSEARCH ）可以一次取多个例如：SEARCH,CAT
	_searchType string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5）可多选，以逗号分隔
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
}

// New
