package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagLogisticsTrajectory 大包物流轨迹查询
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
func CainiaoGlobalImPickupBigbagLogisticsTrajectory(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
