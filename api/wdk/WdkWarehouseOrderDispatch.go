package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWarehouseOrderDispatch 仓作业下发
// wdk.warehouse.order.dispatch
//
// 牵牛花仓作业下发接口提供
func WdkWarehouseOrderDispatch(clt *core.SDKClient, req *wdk.WdkWarehouseOrderDispatchAPIRequest, session string) (*wdk.WdkWarehouseOrderDispatchAPIResponse, error) {
	var resp wdk.WdkWarehouseOrderDispatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
