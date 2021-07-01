package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytSearchbillAPIRequest
通过时间段批量查询入出库单信息 API请求
alibaba.alihealth.drug.kyt.searchbill

通过时间段批量查询入出库单信息 */
type AlibabaAlihealthDrugKytSearchbillAPIRequest struct {
	model.Params
	// 企业标识
	_refEntId string
	// 货主
	_authRefUserId string
	// 开始日期
	_beginDate string
	// 结束日期
	_endDate string
	// 发货企业
	_partnerIdSend string
	// 收货企业
	_partnerIdRecv string
	// 代理企业
	_agentRefUserId string
	// 当前页
	_curPage int64
	// 页大小
	_pageSize int64
	// 单据号码
	_billCode string
	// 单据类型  A : 所有  AI :入库    AO:出库
	_billType string
}

// New
