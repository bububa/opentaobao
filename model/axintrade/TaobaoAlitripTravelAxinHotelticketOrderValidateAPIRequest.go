package axintrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest 阿信度假业务交易试单接口 API请求
// taobao.alitrip.travel.axin.hotelticket.order.validate
//
// 阿信度假业务交易试单接口
type TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest struct {
	model.Params
	// 出行日期
	_serviceStartTime string
	// 门票游玩日期 只能选择出行和结束之间日期
	_useTime string
	// 结束日期
	_serviceEndTime string
	// 产品线
	_bizLine string
	// 产品ID
	_productId int64
	// 购买数量
	_buyAmount int64
	// 分销商ID
	_distributorTid int64
	// 产品单价，单位为分
	_productPrice int64
	// 订单总金额，单位为分
	_totalPrice int64
	// 联系人信息
	_contactInfo *TravelerDto
	// 出行人信息
	_travelerInfoList *TravelerDto
}

// NewTaobaoalitriptravelaxinhotelticketordervalidateRequest 初始化TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest对象
func NewTaobaoalitriptravelaxinhotelticketordervalidateRequest() *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest {
	return &TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.axin.hotelticket.order.validate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceStartTime is ServiceStartTime Setter
// 出行日期
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetServiceStartTime(_serviceStartTime string) error {
	r._serviceStartTime = _serviceStartTime
	r.Set("service_start_time", _serviceStartTime)
	return nil
}

// GetServiceStartTime ServiceStartTime Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetServiceStartTime() string {
	return r._serviceStartTime
}

// SetUseTime is UseTime Setter
// 门票游玩日期 只能选择出行和结束之间日期
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetUseTime(_useTime string) error {
	r._useTime = _useTime
	r.Set("use_time", _useTime)
	return nil
}

// GetUseTime UseTime Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetUseTime() string {
	return r._useTime
}

// SetServiceEndTime is ServiceEndTime Setter
// 结束日期
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetServiceEndTime(_serviceEndTime string) error {
	r._serviceEndTime = _serviceEndTime
	r.Set("service_end_time", _serviceEndTime)
	return nil
}

// GetServiceEndTime ServiceEndTime Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetServiceEndTime() string {
	return r._serviceEndTime
}

// SetBizLine is BizLine Setter
// 产品线
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetBizLine(_bizLine string) error {
	r._bizLine = _bizLine
	r.Set("biz_line", _bizLine)
	return nil
}

// GetBizLine BizLine Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetBizLine() string {
	return r._bizLine
}

// SetProductId is ProductId Setter
// 产品ID
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetBuyAmount is BuyAmount Setter
// 购买数量
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetBuyAmount(_buyAmount int64) error {
	r._buyAmount = _buyAmount
	r.Set("buy_amount", _buyAmount)
	return nil
}

// GetBuyAmount BuyAmount Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetBuyAmount() int64 {
	return r._buyAmount
}

// SetDistributorTid is DistributorTid Setter
// 分销商ID
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetDistributorTid(_distributorTid int64) error {
	r._distributorTid = _distributorTid
	r.Set("distributor_tid", _distributorTid)
	return nil
}

// GetDistributorTid DistributorTid Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetDistributorTid() int64 {
	return r._distributorTid
}

// SetProductPrice is ProductPrice Setter
// 产品单价，单位为分
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetProductPrice(_productPrice int64) error {
	r._productPrice = _productPrice
	r.Set("product_price", _productPrice)
	return nil
}

// GetProductPrice ProductPrice Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetProductPrice() int64 {
	return r._productPrice
}

// SetTotalPrice is TotalPrice Setter
// 订单总金额，单位为分
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetTotalPrice(_totalPrice int64) error {
	r._totalPrice = _totalPrice
	r.Set("total_price", _totalPrice)
	return nil
}

// GetTotalPrice TotalPrice Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetTotalPrice() int64 {
	return r._totalPrice
}

// SetContactInfo is ContactInfo Setter
// 联系人信息
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetContactInfo(_contactInfo *TravelerDto) error {
	r._contactInfo = _contactInfo
	r.Set("contact_info", _contactInfo)
	return nil
}

// GetContactInfo ContactInfo Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetContactInfo() *TravelerDto {
	return r._contactInfo
}

// SetTravelerInfoList is TravelerInfoList Setter
// 出行人信息
func (r *TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) SetTravelerInfoList(_travelerInfoList *TravelerDto) error {
	r._travelerInfoList = _travelerInfoList
	r.Set("traveler_info_list", _travelerInfoList)
	return nil
}

// GetTravelerInfoList TravelerInfoList Getter
func (r TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest) GetTravelerInfoList() *TravelerDto {
	return r._travelerInfoList
}
