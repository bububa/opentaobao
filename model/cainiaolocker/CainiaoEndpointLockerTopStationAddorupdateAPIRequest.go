package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoEndpointLockerTopStationAddorupdateAPIRequest
增加更新代收点 API请求
cainiao.endpoint.locker.top.station.addorupdate

新增或者修改代收点相关信息 */
type CainiaoEndpointLockerTopStationAddorupdateAPIRequest struct {
	model.Params
	// 站点信息
	_stationInfo *StationInfo
}

// New
