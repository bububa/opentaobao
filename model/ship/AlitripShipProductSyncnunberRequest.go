package ship

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
船票班次变更回调 API请求
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

// 初始化AlitripShipProductSyncnunberRequest对象
func NewAlitripShipProductSyncnunberRequest() *AlitripShipProductSyncnunberRequest{
    return &AlitripShipProductSyncnunberRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripShipProductSyncnunberRequest) GetApiMethodName() string {
    return "alitrip.ship.product.syncnunber"
}

// IRequest interface 方法, 获取API参数
func (r AlitripShipProductSyncnunberRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CityName Setter
// 出发城市
func (r *AlitripShipProductSyncnunberRequest) SetCityName(cityName string) error {
    r.cityName = cityName
    r.Set("city_name", cityName)
    return nil
}

// CityName Getter
func (r AlitripShipProductSyncnunberRequest) GetCityName() string {
    return r.cityName
}
// CityCode Setter
// 出发城市code
func (r *AlitripShipProductSyncnunberRequest) SetCityCode(cityCode string) error {
    r.cityCode = cityCode
    r.Set("city_code", cityCode)
    return nil
}

// CityCode Getter
func (r AlitripShipProductSyncnunberRequest) GetCityCode() string {
    return r.cityCode
}
// FromStationName Setter
// 出发港口
func (r *AlitripShipProductSyncnunberRequest) SetFromStationName(fromStationName string) error {
    r.fromStationName = fromStationName
    r.Set("from_station_name", fromStationName)
    return nil
}

// FromStationName Getter
func (r AlitripShipProductSyncnunberRequest) GetFromStationName() string {
    return r.fromStationName
}
// FromStationCode Setter
// 出发港口编号
func (r *AlitripShipProductSyncnunberRequest) SetFromStationCode(fromStationCode string) error {
    r.fromStationCode = fromStationCode
    r.Set("from_station_code", fromStationCode)
    return nil
}

// FromStationCode Getter
func (r AlitripShipProductSyncnunberRequest) GetFromStationCode() string {
    return r.fromStationCode
}
