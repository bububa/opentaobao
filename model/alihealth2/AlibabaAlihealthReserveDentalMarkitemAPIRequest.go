package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalMarkitemAPIRequest
标记商品是否可预约 API请求
alibaba.alihealth.reserve.dental.markitem

标记商品是否可预约 */
type AlibabaAlihealthReserveDentalMarkitemAPIRequest struct {
	model.Params
	// 平台商品id
	_itemId int64
	// 是否可预约，1.可预约 0.不可预约
	_status int64
}

// New
