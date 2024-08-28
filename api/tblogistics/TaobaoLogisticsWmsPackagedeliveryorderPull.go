package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackagedeliveryorderPull 包裹出库单拉单
// taobao.logistics.wms.packagedeliveryorder.pull
//
// 包裹出库单拉单
func TaobaoLogisticsWmsPackagedeliveryorderPull(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackagedeliveryorderPullAPIRequest, resp *tblogistics.TaobaoLogisticsWmsPackagedeliveryorderPullAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
