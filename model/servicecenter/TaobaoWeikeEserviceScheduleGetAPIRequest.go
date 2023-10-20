package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceScheduleGetAPIRequest 客服排班信息查询接口 API请求
// taobao.weike.eservice.schedule.get
//
// 客服排班信息查询接口
type TaobaoWeikeEserviceScheduleGetAPIRequest struct {
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

// NewTaobaoWeikeEserviceScheduleGetRequest 初始化TaobaoWeikeEserviceScheduleGetAPIRequest对象
func NewTaobaoWeikeEserviceScheduleGetRequest() *TaobaoWeikeEserviceScheduleGetAPIRequest {
	return &TaobaoWeikeEserviceScheduleGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) Reset() {
	r._sellerNick = ""
	r._spNick = ""
	r._startDate = ""
	r._endDate = ""
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.schedule.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSellerNick is SellerNick Setter
// 商家子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetSpNick is SpNick Setter
// 服务商子账号昵称，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetSpNick(_spNick string) error {
	r._spNick = _spNick
	r.Set("sp_nick", _spNick)
	return nil
}

// GetSpNick SpNick Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetSpNick() string {
	return r._spNick
}

// SetStartDate is StartDate Setter
// 起始日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束日期，起始日期和结束日期跨度不能超过31天
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetOrderId is OrderId Setter
// 订单ID，orderId、sellerNick、spNick三者不能同时为Null
func (r *TaobaoWeikeEserviceScheduleGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoWeikeEserviceScheduleGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoWeikeEserviceScheduleGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeikeEserviceScheduleGetRequest()
	},
}

// GetTaobaoWeikeEserviceScheduleGetRequest 从 sync.Pool 获取 TaobaoWeikeEserviceScheduleGetAPIRequest
func GetTaobaoWeikeEserviceScheduleGetAPIRequest() *TaobaoWeikeEserviceScheduleGetAPIRequest {
	return poolTaobaoWeikeEserviceScheduleGetAPIRequest.Get().(*TaobaoWeikeEserviceScheduleGetAPIRequest)
}

// ReleaseTaobaoWeikeEserviceScheduleGetAPIRequest 将 TaobaoWeikeEserviceScheduleGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeikeEserviceScheduleGetAPIRequest(v *TaobaoWeikeEserviceScheduleGetAPIRequest) {
	v.Reset()
	poolTaobaoWeikeEserviceScheduleGetAPIRequest.Put(v)
}
