package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest
推广组下的创意报表效果数据查询(汇总数据，不分类型) API请求
taobao.simba.rpt.adgroupcreativeeffect.get

推广组下的创意报表效果数据查询(汇总数据，不分类型) */
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIRequest struct {
	model.Params
	// 权限验证信息
	_subwayToken string
	// 昵称
	_nick string
	// 开始日期，格式yyyy-mm-dd
	_startTime string
	// 结束日期，格式yyyy-mm-dd
	_endTime string
	// 查询推广计划id
	_campaignId int64
	// 推广组id
	_adgroupId int64
	// 数据来源（PC站内：1，PC站外：2，无线站内：4，无线站外 : 5，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 报表类型（搜索：SEARCH,类目出价：CAT, 定向投放：NOSEARCH汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如：SEARCH,CAT
	_searchType string
}

// New
