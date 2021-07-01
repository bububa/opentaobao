package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthReserveDentalBindshopanditemAPIRequest
绑定门店信息，商品信息 API请求
alibaba.alihealth.reserve.dental.bindshopanditem

绑定门店信息，商品信息 */
type AlibabaAlihealthReserveDentalBindshopanditemAPIRequest struct {
	model.Params
	// bind_list
	_bindList []BindDto
}

// New
