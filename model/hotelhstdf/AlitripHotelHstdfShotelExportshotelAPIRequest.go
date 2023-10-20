package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelexportshotelAPIRequest 商家自主导出相似度高的标准酒店 API请求
// alitrip.hotel.hstdf.shotel.exportshotel
//
// 商家通过给出自己的卖家酒店信息，通过接口可以返回相似度高的标准酒店信息
type AlitriphotelhstdfshotelexportshotelAPIRequest struct {
	model.Params
	// 酒店名称，必填
	_name string
	// 电话
	_telNumber string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 地址
	_address string
	// HID，卖家酒店上传到平台后的ID
	_hid int64
	// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
	_cityCode int64
}

// NewAlitriphotelhstdfshotelexportshotelRequest 初始化AlitriphotelhstdfshotelexportshotelAPIRequest对象
func NewAlitriphotelhstdfshotelexportshotelRequest() *AlitriphotelhstdfshotelexportshotelAPIRequest {
	return &AlitriphotelhstdfshotelexportshotelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.exportshotel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 酒店名称，必填
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetName() string {
	return r._name
}

// SetTelNumber is TelNumber Setter
// 电话
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetTelNumber(_telNumber string) error {
	r._telNumber = _telNumber
	r.Set("tel_number", _telNumber)
	return nil
}

// GetTelNumber TelNumber Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetTelNumber() string {
	return r._telNumber
}

// SetLongitude is Longitude Setter
// 经度
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetAddress is Address Setter
// 地址
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetAddress() string {
	return r._address
}

// SetHid is Hid Setter
// HID，卖家酒店上传到平台后的ID
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetHid() int64 {
	return r._hid
}

// SetCityCode is CityCode Setter
// 酒店所在行政区划，对应平台ID，为空时会使用经纬度来定位
func (r *AlitriphotelhstdfshotelexportshotelAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r AlitriphotelhstdfshotelexportshotelAPIRequest) GetCityCode() int64 {
	return r._cityCode
}
