package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikeeserviceschedulegetAPIRequest 客服排班信息查询接口 API请求
// taobao.weike.eservice.schedule.get
//
// 客服排班信息查询接口
type TaobaoweikeeserviceschedulegetAPIRequest struct {
	model.Params
	// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
	_sellerNick string
	// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
	_spNick string
	// 起始日期，起始日期和结束日期跨度不能超过31天
	_startDate string
	// 结束日期，起始日期和结束日期跨度不能超过31天
	_endDate string
	// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
	_orderId int64
}

// NewTaobaoweikeeserviceschedulegetRequest 初始化TaobaoweikeeserviceschedulegetAPIRequest对象
func NewTaobaoweikeeserviceschedulegetRequest() *TaobaoweikeeserviceschedulegetAPIRequest {
	return &TaobaoweikeeserviceschedulegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.schedule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoweikeeserviceschedulegetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetSpNick is SpNick Setter
// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoweikeeserviceschedulegetAPIRequest) SetSpNick(_spNick string) error {
	r._spNick = _spNick
	r.Set("sp_nick", _spNick)
	return nil
}

// GetSpNick SpNick Getter
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetSpNick() string {
	return r._spNick
}

// SetStartDate is StartDate Setter
// 起始日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoweikeeserviceschedulegetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoweikeeserviceschedulegetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetOrderId is OrderId Setter
// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoweikeeserviceschedulegetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoweikeeserviceschedulegetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
