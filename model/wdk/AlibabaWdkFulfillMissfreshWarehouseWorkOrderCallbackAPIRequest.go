package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest
每日优鲜仓作业单回传接口 API请求
alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback

家乐福仓作业单回传接口 */
type AlibabaWdkFulfillMissfreshWarehouseWorkOrderCallbackAPIRequest struct {
	model.Params
	// 作业单回传对象
	_callbackOrder *MissfreshO2OCallbackOrder
}

// New
