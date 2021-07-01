package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest
标准化仓作业单回传接口 API请求
alibaba.wdk.fulfill.warehouse.work.order.callback

标准化仓作业单回传接口 */
type AlibabaWdkFulfillWarehouseWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DrfHalfDayCcCallbackOrder
}

// New
