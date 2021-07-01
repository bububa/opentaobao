package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest
零售药店查询本企业上游企业出库单据信息 API请求
alibaba.alihealth.drugtrace.top.lsyd.listupout

查询货主/本企业上游企业出库单据信息 */
type AlibabaAlihealthDrugtraceTopLsydListupoutAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期（不写时分秒）
	_beginDate string
	// 结束日期（不写时分秒）
	_endDate string
	// 生产批号
	_produceBatchNo string
	// 药品ID
	_drugEntBaseInfoId string
	// 单据类型
	_billType string
	// 药品类型
	_physicType string
	// 状态
	_status string
	// 单据号
	_billCode string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
	// 发货单位
	_fromUserId string
}

// New
