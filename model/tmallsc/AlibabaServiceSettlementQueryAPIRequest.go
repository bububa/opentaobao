package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServiceSettlementQueryAPIRequest
服务平台结算单明细查询服务 API请求
alibaba.service.settlement.query

给服务商提供结算单明细查询功能 */
type AlibabaServiceSettlementQueryAPIRequest struct {
	model.Params
	// 账单查询开始时间。格式示例 2019-03-26 17:15:28
	_gmtCreateStart string
	// 账单查询结束时间，时间区间限制未15分钟。 格式示例 2019-03-26 17:15:28
	_gmtCreateEnd string
	// 当前页面，开始值为1
	_currentPage int64
	// 页面展示条数大小
	_pageSize int64
	// 工单ID
	_workcardId int64
	// 交易主订单号码
	_parentTradeOrderId int64
	// 服务单号
	_serviceOrderId int64
	// 交易实物订单号
	_masterTradeOrderId int64
	// 交易服务订单号
	_serviceTradeOrderId int64
	// 账单修改开始时间。
	_gmtModifiedEnd string
	// 账单修改结束时间，时间区间限制未15分钟。
	_gmtModifiedStart string
}

// New
