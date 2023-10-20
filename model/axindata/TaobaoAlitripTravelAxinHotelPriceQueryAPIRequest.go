package axindata

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelpricequeryAPIRequest 阿信酒店分销-实时报价查询 API请求
// taobao.alitrip.travel.axin.hotel.price.query
//
// 酒店实时报价查询
type TaobaoalitriptravelaxinhotelpricequeryAPIRequest struct {
	model.Params
	// 资源渠道
	_resourceChannel string
	// 入住日期
	_checkIn string
	// 离店日期
	_checkOut string
	// 标准酒店id
	_shid int64
	// 城市code
	_cityCode int64
	// 分销商id
	_distributorTid int64
	// 入住人数
	_nop int64
}

// NewTaobaoalitriptravelaxinhotelpricequeryRequest 初始化TaobaoalitriptravelaxinhotelpricequeryAPIRequest对象
func NewTaobaoalitriptravelaxinhotelpricequeryRequest() *TaobaoalitriptravelaxinhotelpricequeryAPIRequest {
	return &TaobaoalitriptravelaxinhotelpricequeryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.price.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetShid() int64 {
	return r._shid
}

// SetCityCode is CityCode Setter
// 城市code
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetNop is Nop Setter
// 入住人数
func (r *TaobaoalitriptravelaxinhotelpricequeryAPIRequest) SetNop(_nop int64) error {
	r._nop = _nop
	r.Set("nop", _nop)
	return nil
}

// GetNop Nop Getter
func (r TaobaoalitriptravelaxinhotelpricequeryAPIRequest) GetNop() int64 {
	return r._nop
}
