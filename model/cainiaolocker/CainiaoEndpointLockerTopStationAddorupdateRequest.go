package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加更新代收点 APIRequest
cainiao.endpoint.locker.top.station.addorupdate

新增或者修改代收点相关信息
*/
type CainiaoEndpointLockerTopStationAddorupdateRequest struct {
    model.Params

    // 站点信息
    stationInfo   *StationInfo 

}

func NewCainiaoEndpointLockerTopStationAddorupdateRequest() *CainiaoEndpointLockerTopStationAddorupdateRequest{
    return &CainiaoEndpointLockerTopStationAddorupdateRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoEndpointLockerTopStationAddorupdateRequest) GetApiMethodName() string {
    return "cainiao.endpoint.locker.top.station.addorupdate"
}

func (r CainiaoEndpointLockerTopStationAddorupdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoEndpointLockerTopStationAddorupdateRequest) SetStationInfo(stationInfo *StationInfo) error {
    r.stationInfo = stationInfo
    r.Set("station_info", stationInfo)
    return nil
}

func (r CainiaoEndpointLockerTopStationAddorupdateRequest) GetStationInfo() *StationInfo {
    return r.stationInfo
}

