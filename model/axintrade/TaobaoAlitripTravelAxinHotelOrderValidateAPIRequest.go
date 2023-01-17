package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest 阿信酒店订单数据校验 API请求
// taobao.alitrip.travel.axin.hotel.order.validate
//
// 阿信酒店订单下单前的数据校验，包括酒店、房型、售卖政策、入离日期等参数的校验
type TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest struct {
	model.Params
	// 入住日期
	_checkIn string
	// 离店日期
	_checkOut string
	// 资源渠道
	_resourceChannel string
	// 每间房入住人数
	_numberOfAdultsPerRoom string
	// NORMAL-正常场景，SPECIAL-特殊场景
	_invokeScene string
	// 分销商id
	_distributorTid int64
	// 酒店id
	_shid int64
	// 房型id
	_srid int64
	// 商品id
	_itemId int64
	// rpid
	_ratePlanId int64
	// 房间数
	_roomNumber int64
	// rate_id
	_rateId int64
	// 总价
	_totalFee int64
}

// NewTaobaoAlitripTravelAxinHotelOrderValidateRequest 初始化TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelOrderValidateRequest() *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest {
	return &TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetNumberOfAdultsPerRoom is NumberOfAdultsPerRoom Setter
// 每间房入住人数
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetNumberOfAdultsPerRoom(_numberOfAdultsPerRoom string) error {
	r._numberOfAdultsPerRoom = _numberOfAdultsPerRoom
	r.Set("number_of_adults_per_room", _numberOfAdultsPerRoom)
	return nil
}

// GetNumberOfAdultsPerRoom NumberOfAdultsPerRoom Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetNumberOfAdultsPerRoom() string {
	return r._numberOfAdultsPerRoom
}

// SetInvokeScene is InvokeScene Setter
// NORMAL-正常场景，SPECIAL-特殊场景
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetInvokeScene(_invokeScene string) error {
	r._invokeScene = _invokeScene
	r.Set("invoke_scene", _invokeScene)
	return nil
}

// GetInvokeScene InvokeScene Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetInvokeScene() string {
	return r._invokeScene
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetShid is Shid Setter
// 酒店id
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetShid() int64 {
	return r._shid
}

// SetSrid is Srid Setter
// 房型id
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetSrid() int64 {
	return r._srid
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRatePlanId is RatePlanId Setter
// rpid
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetRatePlanId(_ratePlanId int64) error {
	r._ratePlanId = _ratePlanId
	r.Set("rate_plan_id", _ratePlanId)
	return nil
}

// GetRatePlanId RatePlanId Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetRatePlanId() int64 {
	return r._ratePlanId
}

// SetRoomNumber is RoomNumber Setter
// 房间数
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetRoomNumber(_roomNumber int64) error {
	r._roomNumber = _roomNumber
	r.Set("room_number", _roomNumber)
	return nil
}

// GetRoomNumber RoomNumber Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetRoomNumber() int64 {
	return r._roomNumber
}

// SetRateId is RateId Setter
// rate_id
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetRateId(_rateId int64) error {
	r._rateId = _rateId
	r.Set("rate_id", _rateId)
	return nil
}

// GetRateId RateId Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetRateId() int64 {
	return r._rateId
}

// SetTotalFee is TotalFee Setter
// 总价
func (r *TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) SetTotalFee(_totalFee int64) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r TaobaoAlitripTravelAxinHotelOrderValidateAPIRequest) GetTotalFee() int64 {
	return r._totalFee
}
