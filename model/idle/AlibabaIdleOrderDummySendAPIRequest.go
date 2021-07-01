package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleOrderDummySendAPIRequest
闲鱼无需物流发货 API请求
alibaba.idle.order.dummy.send

适用于电子卡券等虚拟商品不需要物流的商品发货。 */
type AlibabaIdleOrderDummySendAPIRequest struct {
	model.Params
	// 请求
	_paramOrderDummySendRequest *OrderDummySendRequest
}

// New
