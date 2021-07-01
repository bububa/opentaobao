package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDentalItemUnbindAPIRequest
ISV解绑商品 API请求
alibaba.alihealth.dental.item.unbind

ISV解绑商品 */
type AlibabaAlihealthDentalItemUnbindAPIRequest struct {
	model.Params
	// ISV门店ID
	_spStoreId string
	// ISV商品ID
	_spItemId string
}

// New
