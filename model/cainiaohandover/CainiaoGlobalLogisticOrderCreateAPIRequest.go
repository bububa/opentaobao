package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalLogisticOrderCreateAPIRequest
创建物流订单 API请求
cainiao.global.logistic.order.create

创建物流订单 */
type CainiaoGlobalLogisticOrderCreateAPIRequest struct {
	model.Params
	// 订单参数
	_orderParam *OpenOrderParam
	// 多语言
	_locale string
}

// New
