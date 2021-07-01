package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaRptCustbaseGetAPIRequest
客户账户报表基础数据对象 API请求
taobao.simba.rpt.custbase.get

客户账户报表基础数据对象 */
type TaobaoSimbaRptCustbaseGetAPIRequest struct {
	model.Params
	// 权限验证信息
	_subwayToken string
	// 昵称
	_nick string
	// 开始日期，格式yyyy-mm-dd
	_startTime string
	// 结束日期，格式yyyy-mm-dd
	_endTime string
	// 页码
	_pageNo int64
	// 每页大小
	_pageSize int64
	// 数据来源（站内：1，站外：2 ，汇总：SUMMARY）SUMMARY必须单选，其他值可多选例如1,2
	_source string
}

// New
