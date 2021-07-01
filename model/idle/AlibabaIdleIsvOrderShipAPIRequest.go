package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderShipAPIRequest
闲鱼无忧购服务商发货 API请求
alibaba.idle.isv.order.ship

闲鱼无忧购业务入仓模式下服务商订单发货的接口 */
type AlibabaIdleIsvOrderShipAPIRequest struct {
	model.Params
	// 订单号
	_bizOrderId string
	// 物流公司
	_logisticsCompany string
	// 运单号
	_shipMailNo string
}

// New
