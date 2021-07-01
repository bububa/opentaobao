package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest
退仓结果回传 API请求
alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed

退货入仓结果回传 */
type AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest struct {
	model.Params
	// 退仓结果
	_returnWarehouseResult *ReturnWarehouseResult
}

// New
