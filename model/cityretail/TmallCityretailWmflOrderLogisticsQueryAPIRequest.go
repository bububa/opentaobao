package cityretail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCityretailWmflOrderLogisticsQueryAPIRequest
完美履约订单物流状态查询接口 API请求
tmall.cityretail.wmfl.order.logistics.query

完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单 */
type TmallCityretailWmflOrderLogisticsQueryAPIRequest struct {
	model.Params
	// 订单号
	_mainOrderId string
}

// New
