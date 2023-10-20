package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxipricegetAPIRequest 获取价格预估信息 API请求
// alibaba.happytrip.taxi.price.get
//
// 打车价格预估
type AlibabahappytriptaxipricegetAPIRequest struct {
	model.Params
	// 出发地纬度
	_flat string
	// 出发地经度
	_flng string
	// 目的地纬度
	_tlat string
	// 目的地经度
	_tlng string
	// 地图类型:amap：高德，默认高德地图
	_mapType string
	// 出发城市id
	_city string
	// 预约单必须传（格式例如：2015-06-16 12:00:09）
	_departureTime string
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
	// 供应商车型代码
	_requireLevel string
	// 用户唯一标识
	_uid string
	// 乘客手机号
	_passengerPhone string
	// 0:实时单 1:预约单
	_type int64
	// 0：不拼车 1:允许拼车，默认不拼车
	_carpoolType int64
	// 乘车人数
	_passengerNumber int64
}

// NewAlibabahappytriptaxipricegetRequest 初始化AlibabahappytriptaxipricegetAPIRequest对象
func NewAlibabahappytriptaxipricegetRequest() *AlibabahappytriptaxipricegetAPIRequest {
	return &AlibabahappytriptaxipricegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxipricegetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.price.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxipricegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxipricegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFlat is Flat Setter
// 出发地纬度
func (r *AlibabahappytriptaxipricegetAPIRequest) SetFlat(_flat string) error {
	r._flat = _flat
	r.Set("flat", _flat)
	return nil
}

// GetFlat Flat Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetFlat() string {
	return r._flat
}

// SetFlng is Flng Setter
// 出发地经度
func (r *AlibabahappytriptaxipricegetAPIRequest) SetFlng(_flng string) error {
	r._flng = _flng
	r.Set("flng", _flng)
	return nil
}

// GetFlng Flng Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetFlng() string {
	return r._flng
}

// SetTlat is Tlat Setter
// 目的地纬度
func (r *AlibabahappytriptaxipricegetAPIRequest) SetTlat(_tlat string) error {
	r._tlat = _tlat
	r.Set("tlat", _tlat)
	return nil
}

// GetTlat Tlat Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetTlat() string {
	return r._tlat
}

// SetTlng is Tlng Setter
// 目的地经度
func (r *AlibabahappytriptaxipricegetAPIRequest) SetTlng(_tlng string) error {
	r._tlng = _tlng
	r.Set("tlng", _tlng)
	return nil
}

// GetTlng Tlng Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetTlng() string {
	return r._tlng
}

// SetMapType is MapType Setter
// 地图类型:amap：高德，默认高德地图
func (r *AlibabahappytriptaxipricegetAPIRequest) SetMapType(_mapType string) error {
	r._mapType = _mapType
	r.Set("map_type", _mapType)
	return nil
}

// GetMapType MapType Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetMapType() string {
	return r._mapType
}

// SetCity is City Setter
// 出发城市id
func (r *AlibabahappytriptaxipricegetAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetCity() string {
	return r._city
}

// SetDepartureTime is DepartureTime Setter
// 预约单必须传（格式例如：2015-06-16 12:00:09）
func (r *AlibabahappytriptaxipricegetAPIRequest) SetDepartureTime(_departureTime string) error {
	r._departureTime = _departureTime
	r.Set("departure_time", _departureTime)
	return nil
}

// GetDepartureTime DepartureTime Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetDepartureTime() string {
	return r._departureTime
}

// SetCostCenter is CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabahappytriptaxipricegetAPIRequest) SetCostCenter(_costCenter string) error {
	r._costCenter = _costCenter
	r.Set("cost_center", _costCenter)
	return nil
}

// GetCostCenter CostCenter Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetCostCenter() string {
	return r._costCenter
}

// SetRequireLevel is RequireLevel Setter
// 供应商车型代码
func (r *AlibabahappytriptaxipricegetAPIRequest) SetRequireLevel(_requireLevel string) error {
	r._requireLevel = _requireLevel
	r.Set("require_level", _requireLevel)
	return nil
}

// GetRequireLevel RequireLevel Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetRequireLevel() string {
	return r._requireLevel
}

// SetUid is Uid Setter
// 用户唯一标识
func (r *AlibabahappytriptaxipricegetAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetUid() string {
	return r._uid
}

// SetPassengerPhone is PassengerPhone Setter
// 乘客手机号
func (r *AlibabahappytriptaxipricegetAPIRequest) SetPassengerPhone(_passengerPhone string) error {
	r._passengerPhone = _passengerPhone
	r.Set("passenger_phone", _passengerPhone)
	return nil
}

// GetPassengerPhone PassengerPhone Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetPassengerPhone() string {
	return r._passengerPhone
}

// SetType is Type Setter
// 0:实时单 1:预约单
func (r *AlibabahappytriptaxipricegetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetType() int64 {
	return r._type
}

// SetCarpoolType is CarpoolType Setter
// 0：不拼车 1:允许拼车，默认不拼车
func (r *AlibabahappytriptaxipricegetAPIRequest) SetCarpoolType(_carpoolType int64) error {
	r._carpoolType = _carpoolType
	r.Set("carpool_type", _carpoolType)
	return nil
}

// GetCarpoolType CarpoolType Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetCarpoolType() int64 {
	return r._carpoolType
}

// SetPassengerNumber is PassengerNumber Setter
// 乘车人数
func (r *AlibabahappytriptaxipricegetAPIRequest) SetPassengerNumber(_passengerNumber int64) error {
	r._passengerNumber = _passengerNumber
	r.Set("passenger_number", _passengerNumber)
	return nil
}

// GetPassengerNumber PassengerNumber Getter
func (r AlibabahappytriptaxipricegetAPIRequest) GetPassengerNumber() int64 {
	return r._passengerNumber
}
