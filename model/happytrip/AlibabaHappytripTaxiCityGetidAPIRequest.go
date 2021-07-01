package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiCityGetidAPIRequest
获取城市id API请求
alibaba.happytrip.taxi.city.getid

通过经纬度坐标返回城市id */
type AlibabaHappytripTaxiCityGetidAPIRequest struct {
	model.Params
	// 纬度
	_lat string
	// 经度
	_lng string
	// 地图类型:amap：高德，默认高德地图
	_mapType string
}

// NewAlibabaHappytripTaxiCityGetidRequest 初始化AlibabaHappytripTaxiCityGetidAPIRequest对象
func NewAlibabaHappytripTaxiCityGetidRequest() *AlibabaHappytripTaxiCityGetidAPIRequest {
	return &AlibabaHappytripTaxiCityGetidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiCityGetidAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.city.getid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiCityGetidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Lat Setter
// 纬度
func (r *AlibabaHappytripTaxiCityGetidAPIRequest) SetLat(_lat string) error {
	r._lat = _lat
	r.Set("lat", _lat)
	return nil
}

// Get Lat Getter
func (r AlibabaHappytripTaxiCityGetidAPIRequest) GetLat() string {
	return r._lat
}

// Set is Lng Setter
// 经度
func (r *AlibabaHappytripTaxiCityGetidAPIRequest) SetLng(_lng string) error {
	r._lng = _lng
	r.Set("lng", _lng)
	return nil
}

// Get Lng Getter
func (r AlibabaHappytripTaxiCityGetidAPIRequest) GetLng() string {
	return r._lng
}

// Set is MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabaHappytripTaxiCityGetidAPIRequest) SetMapType(_mapType string) error {
	r._mapType = _mapType
	r.Set("map_type", _mapType)
	return nil
}

// Get MapType Getter
func (r AlibabaHappytripTaxiCityGetidAPIRequest) GetMapType() string {
	return r._mapType
}
