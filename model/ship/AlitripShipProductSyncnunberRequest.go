package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
船票班次变更回调 APIRequest
alitrip.ship.product.syncnunber

船票班次变更回调
*/
type AlitripShipProductSyncnunberRequest struct {
    model.Params

    // 出发城市
    cityName   string 

    // 出发城市code
    cityCode   string 

    // 出发港口
    fromStationName   string 

    // 出发港口编号
    fromStationCode   string 

}

func NewAlitripShipProductSyncnunberRequest() *AlitripShipProductSyncnunberRequest{
    return &AlitripShipProductSyncnunberRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripShipProductSyncnunberRequest) GetApiMethodName() string {
    return "alitrip.ship.product.syncnunber"
}

func (r AlitripShipProductSyncnunberRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripShipProductSyncnunberRequest) SetCityName(cityName string) error {
    r.cityName = cityName
    r.Set("city_name", cityName)
    return nil
}

func (r AlitripShipProductSyncnunberRequest) GetCityName() string {
    return r.cityName
}

func (r *AlitripShipProductSyncnunberRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

func (r AlitripShipProductSyncnunberRequest) GetCityCode() string {
    return r.cityCode
}

func (r *AlitripShipProductSyncnunberRequest) SetFromStationName(fromStationName string) error {
    r.fromStationName = fromStationName
    r.Set("from_station_name", fromStationName)
    return nil
}

func (r AlitripShipProductSyncnunberRequest) GetFromStationName() string {
    return r.fromStationName
}

func (r *AlitripShipProductSyncnunberRequest) SetFromStationCode(fromStationCode string) error {
    r.fromStationCode = fromStationCode
    r.Set("from_station_code", fromStationCode)
    return nil
}

func (r AlitripShipProductSyncnunberRequest) GetFromStationCode() string {
    return r.fromStationCode
}

