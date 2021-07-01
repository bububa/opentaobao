package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytSearchstatusAPIRequest
单据处理状态查询 API请求
alibaba.alihealth.drug.kyt.searchstatus

单据处理状态查询 */
type AlibabaAlihealthDrugKytSearchstatusAPIRequest struct {
	model.Params
	// 企业ref_ent_id（货主企业的ref_ent_id）
	_refEntId string
	// 开始日期（没有时分秒）
	_beginDate string
	// 结束日期（没有时分秒）
	_endDate string
	// 单据类型 A：全部 AI：全部入库 AO：全部出库
	_billType string
	// 单据号（精确值，不支持模糊查询）
	_billCode string
	// 药品类型
	_drugType string
	// 状态  0, 处理中     3, 处理成功     4, 处理失败
	_dealStatus string
	// 发货商
	_fromUserId string
	// 收货商
	_toUserId string
	// 代理商（第三方物流企业）
	_agentRefUserId string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// New
