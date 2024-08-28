package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackagedeliveryorderConfirm 包裹出库单确认
// taobao.logistics.wms.packagedeliveryorder.confirm
//
// 包裹出库单确认
func TaobaoLogisticsWmsPackagedeliveryorderConfirm(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIRequest, resp *tblogistics.TaobaoLogisticsWmsPackagedeliveryorderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
