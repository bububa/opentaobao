package axindata

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest 阿信酒店分销-实时报价查询 API请求
// taobao.alitrip.travel.axin.hotel.price.query
//
// 酒店实时报价查询
type TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest struct {
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

// NewTaobaoAlitripTravelAxinHotelPriceQueryRequest 初始化TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelPriceQueryRequest() *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest {
	return &TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) Reset() {
	r._resourceChannel = ""
	r._checkIn = ""
	r._checkOut = ""
	r._shid = 0
	r._cityCode = 0
	r._distributorTid = 0
	r._nop = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.price.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetShid() int64 {
	return r._shid
}

// SetCityCode is CityCode Setter
// 城市code
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetNop is Nop Setter
// 入住人数
func (r *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) SetNop(_nop int64) error {
	r._nop = _nop
	r.Set("nop", _nop)
	return nil
}

// GetNop Nop Getter
func (r TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) GetNop() int64 {
	return r._nop
}

var poolTaobaoAlitripTravelAxinHotelPriceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelAxinHotelPriceQueryRequest()
	},
}

// GetTaobaoAlitripTravelAxinHotelPriceQueryRequest 从 sync.Pool 获取 TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest
func GetTaobaoAlitripTravelAxinHotelPriceQueryAPIRequest() *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest {
	return poolTaobaoAlitripTravelAxinHotelPriceQueryAPIRequest.Get().(*TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest)
}

// ReleaseTaobaoAlitripTravelAxinHotelPriceQueryAPIRequest 将 TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelAxinHotelPriceQueryAPIRequest(v *TaobaoAlitripTravelAxinHotelPriceQueryAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelAxinHotelPriceQueryAPIRequest.Put(v)
}
