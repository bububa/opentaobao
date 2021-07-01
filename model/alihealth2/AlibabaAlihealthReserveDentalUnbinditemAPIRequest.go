package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalUnbinditemAPIRequest
解绑商品信息 API请求
alibaba.alihealth.reserve.dental.unbinditem

绑定门店信息，商品信息 */
type AlibabaAlihealthReserveDentalUnbinditemAPIRequest struct {
	model.Params
	// 服务商门店id
	_spStoreId string
	// 服务商商品id
	_spItemId string
}

// New
