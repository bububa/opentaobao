package cainiaolocker

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoendpointlockertopstationaddorupdateAPIRequest 增加更新代收点 API请求
// cainiao.endpoint.locker.top.station.addorupdate
//
// 新增或者修改代收点相关信息
type CainiaoendpointlockertopstationaddorupdateAPIRequest struct {
	model.Params
	// 站点信息
	_stationInfo *StationInfo
}

// NewCainiaoendpointlockertopstationaddorupdateRequest 初始化CainiaoendpointlockertopstationaddorupdateAPIRequest对象
func NewCainiaoendpointlockertopstationaddorupdateRequest() *CainiaoendpointlockertopstationaddorupdateAPIRequest {
	return &CainiaoendpointlockertopstationaddorupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoendpointlockertopstationaddorupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.station.addorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoendpointlockertopstationaddorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoendpointlockertopstationaddorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStationInfo is StationInfo Setter
// 站点信息
func (r *CainiaoendpointlockertopstationaddorupdateAPIRequest) SetStationInfo(_stationInfo *StationInfo) error {
	r._stationInfo = _stationInfo
	r.Set("station_info", _stationInfo)
	return nil
}

// GetStationInfo StationInfo Getter
func (r CainiaoendpointlockertopstationaddorupdateAPIRequest) GetStationInfo() *StationInfo {
	return r._stationInfo
}
