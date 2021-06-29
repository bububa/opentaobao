package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加更新代收点 API请求
cainiao.endpoint.locker.top.station.addorupdate

新增或者修改代收点相关信息
*/
type CainiaoEndpointLockerTopStationAddorupdateRequest struct {
    model.Params
    // 站点信息
    stationInfo   *StationInfo
}

// 初始化CainiaoEndpointLockerTopStationAddorupdateRequest对象
func NewCainiaoEndpointLockerTopStationAddorupdateRequest() *CainiaoEndpointLockerTopStationAddorupdateRequest{
    return &CainiaoEndpointLockerTopStationAddorupdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoEndpointLockerTopStationAddorupdateRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.station.addorupdate"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoEndpointLockerTopStationAddorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StationInfo Setter
// 站点信息
func (r *CainiaoEndpointLockerTopStationAddorupdateRequest) SetStationInfo(stationInfo *StationInfo) error {
    r.stationInfo = stationInfo
    r.Set("station_info", stationInfo)
    return nil
}

// StationInfo Getter
func (r CainiaoEndpointLockerTopStationAddorupdateRequest) GetStationInfo() *StationInfo {
    return r.stationInfo
}
