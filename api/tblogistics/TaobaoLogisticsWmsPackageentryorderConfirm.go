package tblogistics

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsWmsPackageentryorderConfirm 包裹入库单确认
// taobao.logistics.wms.packageentryorder.confirm
//
// 包裹入库单确认
func TaobaoLogisticsWmsPackageentryorderConfirm(ctx context.Context, clt *core.SDKClient, req *tblogistics.TaobaoLogisticsWmsPackageentryorderConfirmAPIRequest, resp *tblogistics.TaobaoLogisticsWmsPackageentryorderConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
