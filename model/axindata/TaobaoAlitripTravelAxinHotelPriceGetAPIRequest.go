package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelpricegetAPIRequest 酒店报价服务-阿信 API请求
// taobao.alitrip.travel.axin.hotel.price.get
//
// 酒店报价查询服务
type TaobaoalitriptravelaxinhotelpricegetAPIRequest struct {
	model.Params
	// 入住日期
	_checkIn string
	// 离店日期
	_checkOut string
	// 标准酒店id
	_shid int64
	// 分销商ID
	_distributorTid int64
	// 城市代码
	_cityCode int64
}

// NewTaobaoalitriptravelaxinhotelpricegetRequest 初始化TaobaoalitriptravelaxinhotelpricegetAPIRequest对象
func NewTaobaoalitriptravelaxinhotelpricegetRequest() *TaobaoalitriptravelaxinhotelpricegetAPIRequest {
	return &TaobaoalitriptravelaxinhotelpricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoalitriptravelaxinhotelpricegetAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoalitriptravelaxinhotelpricegetAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoalitriptravelaxinhotelpricegetAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID
func (r *TaobaoalitriptravelaxinhotelpricegetAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetCityCode is CityCode Setter
// 城市代码
func (r *TaobaoalitriptravelaxinhotelpricegetAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoalitriptravelaxinhotelpricegetAPIRequest) GetCityCode() int64 {
	return r._cityCode
}
