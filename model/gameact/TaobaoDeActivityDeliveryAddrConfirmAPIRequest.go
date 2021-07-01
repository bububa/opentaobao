package gameact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDeActivityDeliveryAddrConfirmAPIRequest
用户收件地址确认 API请求
taobao.de.activity.delivery.addr.confirm

用户收件地址确认 */
type TaobaoDeActivityDeliveryAddrConfirmAPIRequest struct {
	model.Params
	// 加密流水号
	_serialNumber string
	// 地址Sign
	_addressSign string
}

// New
