package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRecycleOrderFulfillmentAPIRequest
闲鱼回收订单履约V1 API请求
alibaba.idle.recycle.order.fulfillment

外部回收商针对自有回收订单的履行 */
type AlibabaIdleRecycleOrderFulfillmentAPIRequest struct {
	model.Params
	// 订单同步入参
	_param0 *RecycleOrderSynDto
}

// New
