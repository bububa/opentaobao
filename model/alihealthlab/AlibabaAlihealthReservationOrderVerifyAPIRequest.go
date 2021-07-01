package alihealthlab

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReservationOrderVerifyAPIRequest
预约单核销接口 API请求
alibaba.alihealth.reservation.order.verify

预约单核销 */
type AlibabaAlihealthReservationOrderVerifyAPIRequest struct {
	model.Params
	// 请求参数
	_verify *VerifyOrderRequest
}

// New
