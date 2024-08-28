package cainiaolocker

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

// CainiaoEndpointLockerTopStationAddorupdate 增加更新代收点
// cainiao.endpoint.locker.top.station.addorupdate
//
// 新增或者修改代收点相关信息
func CainiaoEndpointLockerTopStationAddorupdate(ctx context.Context, clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopStationAddorupdateAPIRequest, resp *cainiaolocker.CainiaoEndpointLockerTopStationAddorupdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
