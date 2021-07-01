package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptAdgroupkeywordbaseGetAPIRequest
推广组下的词基础报表数据查询(明细数据不分类型查询) API请求
taobao.simba.rpt.adgroupkeywordbase.get

推广组下的词基础报表数据查询(明细数据不分类型查询) */
type TaobaoSimbaRptAdgroupkeywordbaseGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广计划ID
	_campaignId int64
	// 推广组ID
	_adgroupId int64
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 权限校验参数
	_subwayToken string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH）可多选例如：SEARCH,CAT
	_searchType string
}

// New
