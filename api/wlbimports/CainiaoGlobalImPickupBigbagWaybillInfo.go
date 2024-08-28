package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagWaybillInfo 大包面单查询
// cainiao.global.im.pickup.bigbag.waybill.info
//
// 大包面单查询
func CainiaoGlobalImPickupBigbagWaybillInfo(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagWaybillInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
