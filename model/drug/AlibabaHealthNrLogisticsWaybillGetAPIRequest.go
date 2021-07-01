package drug

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHealthNrLogisticsWaybillGetAPIRequest
电子面单查询接口 API请求
alibaba.health.nr.logistics.waybill.get

商家登录后根据订单号查询物流单号及电子面单信息 */
type AlibabaHealthNrLogisticsWaybillGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// New
