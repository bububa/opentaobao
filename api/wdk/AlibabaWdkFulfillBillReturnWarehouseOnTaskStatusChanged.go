package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChanged 退仓结果回传
// alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed
//
// 退货入仓结果回传
func AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChanged(clt *core.SDKClient, req *wdk.AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIRequest, resp *wdk.AlibabaWdkFulfillBillReturnWarehouseOnTaskStatusChangedAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
