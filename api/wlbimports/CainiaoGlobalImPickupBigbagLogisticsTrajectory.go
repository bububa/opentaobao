package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagLogisticsTrajectory 大包物流轨迹查询
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
func CainiaoGlobalImPickupBigbagLogisticsTrajectory(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
