package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyWarehouseOrderGet 仓作业单获取
// alibaba.tcls.aelophy.warehouse.order.get
//
// 仓作业单获取
func AlibabaTclsAelophyWarehouseOrderGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyWarehouseOrderGetAPIRequest, resp *wdk.AlibabaTclsAelophyWarehouseOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
