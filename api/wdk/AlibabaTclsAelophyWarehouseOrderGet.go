package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyWarehouseOrderGet 仓作业单获取
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
func AlibabaTclsAelophyWarehouseOrderGet(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyWarehouseOrderGetAPIRequest, resp *wdk.AlibabaTclsAelophyWarehouseOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
