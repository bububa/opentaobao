package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest
上传单据后处理状态查询 API请求
alibaba.alihealth.drugtrace.top.lsyd.query.billstatus

单据处理状态查询 */
type AlibabaAlihealthDrugtraceTopLsydQueryBillstatusAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 开始日期
	_beginDate string
	// 结束日期
	_endDate string
	// 单据类型 A：全部 AI：全部入库 AO：全部出库
	_billType string
	// 单据号
	_billCode string
	// 药品类型
	_drugType string
	// 状态  0, 上传成功     3, 处理成功     4, 处理失败
	_dealStatus string
	// 发货商
	_fromUserId string
	// 收货商
	_toUserId string
	// 代理商
	_agentRefUserId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// New
