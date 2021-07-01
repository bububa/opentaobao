package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest
末端配配送作业回传 API请求
alibaba.wdk.fulfill.dms.delivery.work.order.callback

末端配配送作业回传。 */
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DeliveryCallbackOrder
}

// New
