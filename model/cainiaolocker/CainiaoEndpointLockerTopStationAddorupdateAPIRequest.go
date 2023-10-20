package cainiaolocker

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoEndpointLockerTopStationAddorupdateAPIRequest 增加更新代收点 API请求
// cainiao.endpoint.locker.top.station.addorupdate
//
// 新增或者修改代收点相关信息
type CainiaoEndpointLockerTopStationAddorupdateAPIRequest struct {
	model.Params
	// 站点信息
	_stationInfo *StationInfo
}

// NewCainiaoEndpointLockerTopStationAddorupdateRequest 初始化CainiaoEndpointLockerTopStationAddorupdateAPIRequest对象
func NewCainiaoEndpointLockerTopStationAddorupdateRequest() *CainiaoEndpointLockerTopStationAddorupdateAPIRequest {
	return &CainiaoEndpointLockerTopStationAddorupdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoEndpointLockerTopStationAddorupdateAPIRequest) Reset() {
	r._stationInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetApiMethodName() string {
	return "cainiao.endpoint.locker.top.station.addorupdate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStationInfo is StationInfo Setter
// 站点信息
func (r *CainiaoEndpointLockerTopStationAddorupdateAPIRequest) SetStationInfo(_stationInfo *StationInfo) error {
	r._stationInfo = _stationInfo
	r.Set("station_info", _stationInfo)
	return nil
}

// GetStationInfo StationInfo Getter
func (r CainiaoEndpointLockerTopStationAddorupdateAPIRequest) GetStationInfo() *StationInfo {
	return r._stationInfo
}

var poolCainiaoEndpointLockerTopStationAddorupdateAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoEndpointLockerTopStationAddorupdateRequest()
	},
}

// GetCainiaoEndpointLockerTopStationAddorupdateRequest 从 sync.Pool 获取 CainiaoEndpointLockerTopStationAddorupdateAPIRequest
func GetCainiaoEndpointLockerTopStationAddorupdateAPIRequest() *CainiaoEndpointLockerTopStationAddorupdateAPIRequest {
	return poolCainiaoEndpointLockerTopStationAddorupdateAPIRequest.Get().(*CainiaoEndpointLockerTopStationAddorupdateAPIRequest)
}

// ReleaseCainiaoEndpointLockerTopStationAddorupdateAPIRequest 将 CainiaoEndpointLockerTopStationAddorupdateAPIRequest 放入 sync.Pool
func ReleaseCainiaoEndpointLockerTopStationAddorupdateAPIRequest(v *CainiaoEndpointLockerTopStationAddorupdateAPIRequest) {
	v.Reset()
	poolCainiaoEndpointLockerTopStationAddorupdateAPIRequest.Put(v)
}
