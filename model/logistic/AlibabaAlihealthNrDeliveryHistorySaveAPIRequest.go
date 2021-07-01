package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthNrDeliveryHistorySaveAPIRequest
物流信息回传接口 API请求
alibaba.alihealth.nr.delivery.history.save

商家ERP回传物流信息 */
type AlibabaAlihealthNrDeliveryHistorySaveAPIRequest struct {
	model.Params
	// 入参
	_deliveryEvent *DeliveryEventDto
}

// New
