package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkFulfillWarehouseWorkOrderSealbox 仓封箱回告
// alibaba.wdk.fulfill.warehouse.work.order.sealbox
//
// 仓封箱回告箱与包裹的关系
func AlibabaWdkFulfillWarehouseWorkOrderSealbox(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIRequest, resp *wdk.AlibabaWdkFulfillWarehouseWorkOrderSealboxAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
