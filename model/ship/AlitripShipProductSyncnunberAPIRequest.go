package ship

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripShipProductSyncnunberAPIRequest
船票班次变更回调 API请求
alitrip.ship.product.syncnunber

船票班次变更回调 */
type AlitripShipProductSyncnunberAPIRequest struct {
	model.Params
	// 出发城市
	_cityName string
	// 出发城市code
	_cityCode string
	// 出发港口
	_fromStationName string
	// 出发港口编号
	_fromStationCode string
}

// NewAlitripShipProductSyncnunberRequest 初始化AlitripShipProductSyncnunberAPIRequest对象
func NewAlitripShipProductSyncnunberRequest() *AlitripShipProductSyncnunberAPIRequest {
	return &AlitripShipProductSyncnunberAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripShipProductSyncnunberAPIRequest) GetApiMethodName() string {
	return "alitrip.ship.product.syncnunber"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripShipProductSyncnunberAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CityName Setter
// 出发城市
func (r *AlitripShipProductSyncnunberAPIRequest) SetCityName(_cityName string) error {
	r._cityName = _cityName
	r.Set("city_name", _cityName)
	return nil
}

// Get CityName Getter
func (r AlitripShipProductSyncnunberAPIRequest) GetCityName() string {
	return r._cityName
}

// Set is CityCode Setter
// 出发城市code
func (r *AlitripShipProductSyncnunberAPIRequest) SetCityCode(_cityCode string) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// Get CityCode Getter
func (r AlitripShipProductSyncnunberAPIRequest) GetCityCode() string {
	return r._cityCode
}

// Set is FromStationName Setter
// 出发港口
func (r *AlitripShipProductSyncnunberAPIRequest) SetFromStationName(_fromStationName string) error {
	r._fromStationName = _fromStationName
	r.Set("from_station_name", _fromStationName)
	return nil
}

// Get FromStationName Getter
func (r AlitripShipProductSyncnunberAPIRequest) GetFromStationName() string {
	return r._fromStationName
}

// Set is FromStationCode Setter
// 出发港口编号
func (r *AlitripShipProductSyncnunberAPIRequest) SetFromStationCode(_fromStationCode string) error {
	r._fromStationCode = _fromStationCode
	r.Set("from_station_code", _fromStationCode)
	return nil
}

// Get FromStationCode Getter
func (r AlitripShipProductSyncnunberAPIRequest) GetFromStationCode() string {
	return r._fromStationCode
}
