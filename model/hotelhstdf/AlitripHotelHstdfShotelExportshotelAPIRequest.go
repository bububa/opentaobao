package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripHotelHstdfShotelExportshotelAPIRequest
商家自主导出相似度高的标准酒店 API请求
alitrip.hotel.hstdf.shotel.exportshotel

商家通过给出自己的卖家酒店信息，通过接口可以返回相似度高的标准酒店信息 */
type AlitripHotelHstdfShotelExportshotelAPIRequest struct {
	model.Params
	// HID，卖家酒店上传到平台后的ID
	_hid int64
	// 酒店名称，必填
	_name string
	// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
	_cityCode int64
	// 电话
	_telNumber string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 地址
	_address string
}

// NewAlitripHotelHstdfShotelExportshotelRequest 初始化AlitripHotelHstdfShotelExportshotelAPIRequest对象
func NewAlitripHotelHstdfShotelExportshotelRequest() *AlitripHotelHstdfShotelExportshotelAPIRequest {
	return &AlitripHotelHstdfShotelExportshotelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.exportshotel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Hid Setter
// HID，卖家酒店上传到平台后的ID
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// Get Hid Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetHid() int64 {
	return r._hid
}

// Set is Name Setter
// 酒店名称，必填
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetName() string {
	return r._name
}

// Set is CityCode Setter
// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// Get CityCode Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// Set is TelNumber Setter
// 电话
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetTelNumber(_telNumber string) error {
	r._telNumber = _telNumber
	r.Set("tel_number", _telNumber)
	return nil
}

// Get TelNumber Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetTelNumber() string {
	return r._telNumber
}

// Set is Longitude Setter
// 经度
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// Get Longitude Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetLongitude() string {
	return r._longitude
}

// Set is Latitude Setter
// 纬度
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// Get Latitude Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetLatitude() string {
	return r._latitude
}

// Set is Address Setter
// 地址
func (r *AlitripHotelHstdfShotelExportshotelAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r AlitripHotelHstdfShotelExportshotelAPIRequest) GetAddress() string {
	return r._address
}
