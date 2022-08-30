package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest 阿信度假业务创单并支付接口 API请求
// taobao.alitrip.travel.axin.hotelticket.order.createorder
//
// 阿信度假业务创单并支付接口
type TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest struct {
	model.Params
	// 分销商订单号
	_outerOrderId string
	// 出行日期 格式： yyyy-MM-dd
	_serviceStartTime string
	// 结束日期 格式： yyyy-MM-dd
	_serviceEndTime string
	// 门票游玩日期 只能选择出行和结束之间日期  格式： yyyy-MM-dd
	_useTime string
	// 产品线
	_bizLine string
	// 产品id
	_productId int64
	// 联系人信息
	_contactInfo *TravelerDto
	// 出行人信息
	_travelerInfoList *TravelerDto
	// 购买数量
	_buyAmount int64
	// 产品单价
	_productPrice int64
	// 订单总金额，单位为分
	_totalPrice int64
	// 分销商ID(淘系)
	_distributorTid int64
}

// NewTaobaoAlitripTravelAxinHotelticketOrderCreateorderRequest 初始化TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest对象
func NewTaobaoAlitripTravelAxinHotelticketOrderCreateorderRequest() *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest {
	return &TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.order.createorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterOrderId is OuterOrderId Setter
// 分销商订单号
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetOuterOrderId(_outerOrderId string) error {
	r._outerOrderId = _outerOrderId
	r.Set("outer_order_id", _outerOrderId)
	return nil
}

// GetOuterOrderId OuterOrderId Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetOuterOrderId() string {
	return r._outerOrderId
}

// SetServiceStartTime is ServiceStartTime Setter
// 出行日期 格式： yyyy-MM-dd
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetServiceStartTime(_serviceStartTime string) error {
	r._serviceStartTime = _serviceStartTime
	r.Set("service_start_time", _serviceStartTime)
	return nil
}

// GetServiceStartTime ServiceStartTime Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetServiceStartTime() string {
	return r._serviceStartTime
}

// SetServiceEndTime is ServiceEndTime Setter
// 结束日期 格式： yyyy-MM-dd
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetServiceEndTime(_serviceEndTime string) error {
	r._serviceEndTime = _serviceEndTime
	r.Set("service_end_time", _serviceEndTime)
	return nil
}

// GetServiceEndTime ServiceEndTime Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetServiceEndTime() string {
	return r._serviceEndTime
}

// SetUseTime is UseTime Setter
// 门票游玩日期 只能选择出行和结束之间日期  格式： yyyy-MM-dd
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetUseTime(_useTime string) error {
	r._useTime = _useTime
	r.Set("use_time", _useTime)
	return nil
}

// GetUseTime UseTime Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetUseTime() string {
	return r._useTime
}

// SetBizLine is BizLine Setter
// 产品线
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetBizLine(_bizLine string) error {
	r._bizLine = _bizLine
	r.Set("biz_line", _bizLine)
	return nil
}

// GetBizLine BizLine Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetBizLine() string {
	return r._bizLine
}

// SetProductId is ProductId Setter
// 产品id
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetContactInfo is ContactInfo Setter
// 联系人信息
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetContactInfo(_contactInfo *TravelerDto) error {
	r._contactInfo = _contactInfo
	r.Set("contact_info", _contactInfo)
	return nil
}

// GetContactInfo ContactInfo Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetContactInfo() *TravelerDto {
	return r._contactInfo
}

// SetTravelerInfoList is TravelerInfoList Setter
// 出行人信息
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetTravelerInfoList(_travelerInfoList *TravelerDto) error {
	r._travelerInfoList = _travelerInfoList
	r.Set("traveler_info_list", _travelerInfoList)
	return nil
}

// GetTravelerInfoList TravelerInfoList Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetTravelerInfoList() *TravelerDto {
	return r._travelerInfoList
}

// SetBuyAmount is BuyAmount Setter
// 购买数量
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetBuyAmount(_buyAmount int64) error {
	r._buyAmount = _buyAmount
	r.Set("buy_amount", _buyAmount)
	return nil
}

// GetBuyAmount BuyAmount Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetBuyAmount() int64 {
	return r._buyAmount
}

// SetProductPrice is ProductPrice Setter
// 产品单价
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetProductPrice(_productPrice int64) error {
	r._productPrice = _productPrice
	r.Set("product_price", _productPrice)
	return nil
}

// GetProductPrice ProductPrice Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetProductPrice() int64 {
	return r._productPrice
}

// SetTotalPrice is TotalPrice Setter
// 订单总金额，单位为分
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetTotalPrice(_totalPrice int64) error {
	r._totalPrice = _totalPrice
	r.Set("total_price", _totalPrice)
	return nil
}

// GetTotalPrice TotalPrice Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetTotalPrice() int64 {
	return r._totalPrice
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID(淘系)
func (r *TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoAlitripTravelAxinHotelticketOrderCreateorderAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}
