package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Wdkwarehouseorderdispatch 仓作业下发
// wdk.warehouse.order.dispatch
//
// 牵牛花仓作业下发接口提供
func Wdkwarehouseorderdispatch(clt *core.SDKClient, req *wdk.WdkwarehouseorderdispatchAPIRequest, session string) (*wdk.WdkwarehouseorderdispatchAPIResponse, error) {
	var resp wdk.WdkwarehouseorderdispatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
