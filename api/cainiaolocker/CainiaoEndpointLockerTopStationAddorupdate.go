package cainiaolocker

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaolocker"
)

/* CainiaoEndpointLockerTopStationAddorupdate
增加更新代收点
cainiao.endpoint.locker.top.station.addorupdate

新增或者修改代收点相关信息 */
func CainiaoEndpointLockerTopStationAddorupdate(clt *core.SDKClient, req *cainiaolocker.CainiaoEndpointLockerTopStationAddorupdateAPIRequest, session string) (*cainiaolocker.CainiaoEndpointLockerTopStationAddorupdateAPIResponse, error) {
	var resp cainiaolocker.CainiaoEndpointLockerTopStationAddorupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
