package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest
大润发B2C仓作业单回传接口 API请求
alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback

大润发B2C仓作业单回传接口 */
type AlibabaWdkFulfillRtBtocWarehouseWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *DrfB2CCallbackOrder
}

// New
