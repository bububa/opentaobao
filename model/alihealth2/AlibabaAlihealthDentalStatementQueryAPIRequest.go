package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalStatementQueryAPIRequest
ISV查询对账单 API请求
alibaba.alihealth.dental.statement.query

ISV查询对账单 */
type AlibabaAlihealthDentalStatementQueryAPIRequest struct {
	model.Params
	// 订单ID
	_orderId string
	// 结算周期，单位月
	_statementTime string
}

// New
