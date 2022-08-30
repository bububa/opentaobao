package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// WdkWarehouseOrderCancel 仓作业取消下发
// wdk.warehouse.order.cancel
//
// 仓作业取消下发
func WdkWarehouseOrderCancel(clt *core.SDKClient, req *wdk.WdkWarehouseOrderCancelAPIRequest, session string) (*wdk.WdkWarehouseOrderCancelAPIResponse, error) {
	var resp wdk.WdkWarehouseOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
