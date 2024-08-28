package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagLogisticsTrajectory 大包物流轨迹查询
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
func CainiaoGlobalImPickupBigbagLogisticsTrajectory(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
