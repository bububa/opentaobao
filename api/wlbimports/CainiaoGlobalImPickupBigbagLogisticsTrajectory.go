package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupbigbaglogisticstrajectory 大包物流轨迹查询
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
func Cainiaoglobalimpickupbigbaglogisticstrajectory(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupbigbaglogisticstrajectoryAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupbigbaglogisticstrajectoryAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupbigbaglogisticstrajectoryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
