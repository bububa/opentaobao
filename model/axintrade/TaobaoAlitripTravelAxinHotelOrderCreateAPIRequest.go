package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest 酒店分销订单创建服务-阿信 API请求
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
type TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest struct {
	model.Params
	// 入住人信息
	_customers []CustomerDto
	// 备注
	_remark string
	// 离店日期
	_checkOut string
	// 入住日期
	_checkIn string
	// 酒店名称
	_hotelName string
	// 外部订单号
	_outerOrderId string
	// 订单验证返回的key
	_createKey string
	// 币种
	_currencyCode string
	// 资源渠道
	_resourceChannel string
	// 到店时间
	_hotelArrivalTime *HotelArrivalTime
	// 联系人信息
	_hotelContact *ContactInfoApiDto
	// 总房价（单位分）
	_totalRoomPrice int64
	// 房间数
	_roomNumber int64
	// 售卖政策id
	_ratePlanId int64
	// 商品id
	_itemId int64
	// 标准房型ID
	_srid int64
	// 标准酒店id
	_shid int64
	// 分销商ID
	_distributorTid int64
	// 总优惠金额
	_promotionTotalPrice int64
	// 城市代码
	_cityCode int64
	// 每间房入住人数
	_numberOfAdultsPerRoom int64
	// 人民币总价格
	_totalCnyRoomPrice int64
}

// NewTaobaoAlitripTravelAxinHotelOrderCreateRequest 初始化TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelOrderCreateRequest() *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest {
	return &TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotel.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCustomers is Customers Setter
// 入住人信息
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetCustomers(_customers []CustomerDto) error {
	r._customers = _customers
	r.Set("customers", _customers)
	return nil
}

// GetCustomers Customers Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetCustomers() []CustomerDto {
	return r._customers
}

// SetRemark is Remark Setter
// 备注
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetRemark() string {
	return r._remark
}

// SetCheckOut is CheckOut Setter
// 离店日期
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetCheckOut(_checkOut string) error {
	r._checkOut = _checkOut
	r.Set("check_out", _checkOut)
	return nil
}

// GetCheckOut CheckOut Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetCheckOut() string {
	return r._checkOut
}

// SetCheckIn is CheckIn Setter
// 入住日期
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetCheckIn(_checkIn string) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetCheckIn() string {
	return r._checkIn
}

// SetHotelName is HotelName Setter
// 酒店名称
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetHotelName(_hotelName string) error {
	r._hotelName = _hotelName
	r.Set("hotel_name", _hotelName)
	return nil
}

// GetHotelName HotelName Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetHotelName() string {
	return r._hotelName
}

// SetOuterOrderId is OuterOrderId Setter
// 外部订单号
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetOuterOrderId(_outerOrderId string) error {
	r._outerOrderId = _outerOrderId
	r.Set("outer_order_id", _outerOrderId)
	return nil
}

// GetOuterOrderId OuterOrderId Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetOuterOrderId() string {
	return r._outerOrderId
}

// SetCreateKey is CreateKey Setter
// 订单验证返回的key
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetCreateKey(_createKey string) error {
	r._createKey = _createKey
	r.Set("create_key", _createKey)
	return nil
}

// GetCreateKey CreateKey Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetCreateKey() string {
	return r._createKey
}

// SetCurrencyCode is CurrencyCode Setter
// 币种
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetResourceChannel is ResourceChannel Setter
// 资源渠道
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetResourceChannel(_resourceChannel string) error {
	r._resourceChannel = _resourceChannel
	r.Set("resource_channel", _resourceChannel)
	return nil
}

// GetResourceChannel ResourceChannel Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetResourceChannel() string {
	return r._resourceChannel
}

// SetHotelArrivalTime is HotelArrivalTime Setter
// 到店时间
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetHotelArrivalTime(_hotelArrivalTime *HotelArrivalTime) error {
	r._hotelArrivalTime = _hotelArrivalTime
	r.Set("hotel_arrival_time", _hotelArrivalTime)
	return nil
}

// GetHotelArrivalTime HotelArrivalTime Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetHotelArrivalTime() *HotelArrivalTime {
	return r._hotelArrivalTime
}

// SetHotelContact is HotelContact Setter
// 联系人信息
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetHotelContact(_hotelContact *ContactInfoApiDto) error {
	r._hotelContact = _hotelContact
	r.Set("hotel_contact", _hotelContact)
	return nil
}

// GetHotelContact HotelContact Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetHotelContact() *ContactInfoApiDto {
	return r._hotelContact
}

// SetTotalRoomPrice is TotalRoomPrice Setter
// 总房价（单位分）
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetTotalRoomPrice(_totalRoomPrice int64) error {
	r._totalRoomPrice = _totalRoomPrice
	r.Set("total_room_price", _totalRoomPrice)
	return nil
}

// GetTotalRoomPrice TotalRoomPrice Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetTotalRoomPrice() int64 {
	return r._totalRoomPrice
}

// SetRoomNumber is RoomNumber Setter
// 房间数
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetRoomNumber(_roomNumber int64) error {
	r._roomNumber = _roomNumber
	r.Set("room_number", _roomNumber)
	return nil
}

// GetRoomNumber RoomNumber Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetRoomNumber() int64 {
	return r._roomNumber
}

// SetRatePlanId is RatePlanId Setter
// 售卖政策id
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetRatePlanId(_ratePlanId int64) error {
	r._ratePlanId = _ratePlanId
	r.Set("rate_plan_id", _ratePlanId)
	return nil
}

// GetRatePlanId RatePlanId Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetRatePlanId() int64 {
	return r._ratePlanId
}

// SetItemId is ItemId Setter
// 商品id
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSrid is Srid Setter
// 标准房型ID
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetSrid() int64 {
	return r._srid
}

// SetShid is Shid Setter
// 标准酒店id
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetShid() int64 {
	return r._shid
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetPromotionTotalPrice is PromotionTotalPrice Setter
// 总优惠金额
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetPromotionTotalPrice(_promotionTotalPrice int64) error {
	r._promotionTotalPrice = _promotionTotalPrice
	r.Set("promotion_total_price", _promotionTotalPrice)
	return nil
}

// GetPromotionTotalPrice PromotionTotalPrice Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetPromotionTotalPrice() int64 {
	return r._promotionTotalPrice
}

// SetCityCode is CityCode Setter
// 城市代码
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetCityCode(_cityCode int64) error {
	r._cityCode = _cityCode
	r.Set("city_code", _cityCode)
	return nil
}

// GetCityCode CityCode Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetCityCode() int64 {
	return r._cityCode
}

// SetNumberOfAdultsPerRoom is NumberOfAdultsPerRoom Setter
// 每间房入住人数
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetNumberOfAdultsPerRoom(_numberOfAdultsPerRoom int64) error {
	r._numberOfAdultsPerRoom = _numberOfAdultsPerRoom
	r.Set("number_of_adults_per_room", _numberOfAdultsPerRoom)
	return nil
}

// GetNumberOfAdultsPerRoom NumberOfAdultsPerRoom Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetNumberOfAdultsPerRoom() int64 {
	return r._numberOfAdultsPerRoom
}

// SetTotalCnyRoomPrice is TotalCnyRoomPrice Setter
// 人民币总价格
func (r *TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) SetTotalCnyRoomPrice(_totalCnyRoomPrice int64) error {
	r._totalCnyRoomPrice = _totalCnyRoomPrice
	r.Set("total_cny_room_price", _totalCnyRoomPrice)
	return nil
}

// GetTotalCnyRoomPrice TotalCnyRoomPrice Getter
func (r TaobaoAlitripTravelAxinHotelOrderCreateAPIRequest) GetTotalCnyRoomPrice() int64 {
	return r._totalCnyRoomPrice
}
