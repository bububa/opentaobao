package ascpffo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressFulfillmentEventAPIRequest
AE履约事件处理 API请求
aliexpress.fulfillment.event

AE用 履约底层声明发货能力 */
type AliexpressFulfillmentEventAPIRequest struct {
	model.Params
	// 入参对象
	_param *FulfillmentOrderStatusUpdateRequest
}

// New
