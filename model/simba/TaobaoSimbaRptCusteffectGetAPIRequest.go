package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptCusteffectGetAPIRequest
用户账户报表效果数据查询（只有汇总数据，无分类数据） API请求
taobao.simba.rpt.custeffect.get

用户账户报表效果数据查询（只有汇总数据，无分类数据） */
type TaobaoSimbaRptCusteffectGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 开始时间，格式yyyy-mm-dd
	_startTime string
	// 结束时间，格式yyyy-mm-dd
	_endTime string
	// 权限校验参数
	_subwayToken string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
}

// New
