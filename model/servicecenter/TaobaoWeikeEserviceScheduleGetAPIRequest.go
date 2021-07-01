package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeikeEserviceScheduleGetAPIRequest
客服排班信息查询接口 API请求
taobao.weike.eservice.schedule.get

客服排班信息查询接口 */
type TaobaoWeikeEserviceScheduleGetAPIRequest struct {
	model.Params
	// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
	_orderId int64
	// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
	_sellerNick string
	// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
	_spNick string
	// 起始日期，起始日期和结束日期跨度不能超过31天
	_startDate string
	// 结束日期，起始日期和结束日期跨度不能超过31天
	_endDate string
}

// NewTaobaoWeikeEserviceScheduleGetRequest 初始化TaobaoWeikeEserviceScheduleGetAPIRequest对象
func NewTaobaoWeikeEserviceScheduleGetRequest() *TaobaoWeikeEserviceScheduleGetAPIRequest {
	return &TaobaoWeikeEserviceScheduleGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.schedule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is SellerNick Setter
// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// Get SellerNick Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// Set is SpNick Setter
// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetSpNick(_spNick string) error {
	r._spNick = _spNick
	r.Set("sp_nick", _spNick)
	return nil
}

// Get SpNick Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetSpNick() string {
	return r._spNick
}

// Set is StartDate Setter
// 起始日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is EndDate Setter
// 结束日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetEndDate() string {
	return r._endDate
}
