package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthNrCepOrderQueryAPIRequest
订单详情查询接口 API请求
alibaba.health.nr.cep.order.query

订单详情查询接口 */
type AlibabaHealthNrCepOrderQueryAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// New
