package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyWarehouseOrderGet 仓作业单获取
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
func AlibabaTclsAelophyWarehouseOrderGet(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyWarehouseOrderGetAPIRequest, session string) (*wdk.AlibabaTclsAelophyWarehouseOrderGetAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyWarehouseOrderGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
