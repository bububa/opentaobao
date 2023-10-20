package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillbillreturnwarehouseontaskstatuschanged 退仓结果回传
// alibaba.wdk.fulfill.bill.return.warehouse.on.task.status.changed
//
// 退货入仓结果回传
func Alibabawdkfulfillbillreturnwarehouseontaskstatuschanged(clt *core.SDKClient, req *wdk.AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIRequest, session string) (*wdk.AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillbillreturnwarehouseontaskstatuschangedAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
