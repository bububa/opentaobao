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

// New
