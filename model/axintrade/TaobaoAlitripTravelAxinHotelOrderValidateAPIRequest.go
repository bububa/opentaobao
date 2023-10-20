package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelordervalidateAPIRequest 阿信酒店订单数据校验 API请求
// taobao.alitrip.travel.axin.hotel.order.validate
//
// 阿信酒店订单下单前的数据校验，包括酒店、房型、售卖政策、入离日期等参数的校验
type TaobaoalitriptravelaxinhotelordervalidateAPIRequest struct {
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

// NewTaobaoalitriptravelaxinhotelordervalidateRequest 初始化TaobaoalitriptravelaxinhotelordervalidateAPIRequest对象
func NewTaobaoalitriptravelaxinhotelordervalidateRequest() *TaobaoalitriptravelaxinhotelordervalidateAPIRequest {
	return &TaobaoalitriptravelaxinhotelordervalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetNumberOfAdultsPerRoom is NumberOfAdultsPerRoom Setter
// 每间房入住人数
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetNumberOfAdultsPerRoom(_numberOfAdultsPerRoom string) error {
	r._numberOfAdultsPerRoom = _numberOfAdultsPerRoom
	r.Set("number_of_adults_per_room", _numberOfAdultsPerRoom)
	return nil
}

// GetNumberOfAdultsPerRoom NumberOfAdultsPerRoom Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetNumberOfAdultsPerRoom() string {
	return r._numberOfAdultsPerRoom
}

// SetInvokeScene is InvokeScene Setter
// NORMAL-正常场景，SPECIAL-特殊场景
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetInvokeScene(_invokeScene string) error {
	r._invokeScene = _invokeScene
	r.Set("invoke_scene", _invokeScene)
	return nil
}

// GetInvokeScene InvokeScene Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetInvokeScene() string {
	return r._invokeScene
}

// SetDistributorTid is DistributorTid Setter
// 分销商id
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetShid is Shid Setter
// 酒店id
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetShid() int64 {
	return r._shid
}

// SetSrid is Srid Setter
// 房型id
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetSrid() int64 {
	return r._srid
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRatePlanId is RatePlanId Setter
// rpid
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetRatePlanId(_ratePlanId int64) error {
	r._ratePlanId = _ratePlanId
	r.Set("rate_plan_id", _ratePlanId)
	return nil
}

// GetRatePlanId RatePlanId Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetRatePlanId() int64 {
	return r._ratePlanId
}

// SetRoomNumber is RoomNumber Setter
// 房间数
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetRoomNumber(_roomNumber int64) error {
	r._roomNumber = _roomNumber
	r.Set("room_number", _roomNumber)
	return nil
}

// GetRoomNumber RoomNumber Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetRoomNumber() int64 {
	return r._roomNumber
}

// SetRateId is RateId Setter
// rate_id
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetRateId(_rateId int64) error {
	r._rateId = _rateId
	r.Set("rate_id", _rateId)
	return nil
}

// GetRateId RateId Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetRateId() int64 {
	return r._rateId
}

// SetTotalFee is TotalFee Setter
// 总价
func (r *TaobaoalitriptravelaxinhotelordervalidateAPIRequest) SetTotalFee(_totalFee int64) error {
	r._totalFee = _totalFee
	r.Set("total_fee", _totalFee)
	return nil
}

// GetTotalFee TotalFee Getter
func (r TaobaoalitriptravelaxinhotelordervalidateAPIRequest) GetTotalFee() int64 {
	return r._totalFee
}
