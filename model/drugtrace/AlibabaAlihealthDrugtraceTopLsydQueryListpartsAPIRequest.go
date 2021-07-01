package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest
往来单位查询 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.listparts

查询往来单位列表 */
type AlibabaAlihealthDrugtraceTopLsydQueryListpartsAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 企业名称
	_entName string
	// 企业自定义编号
	_refPartnerId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
}

// New
