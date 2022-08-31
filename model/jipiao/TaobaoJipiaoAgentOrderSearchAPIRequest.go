package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJipiaoAgentOrderSearchAPIRequest 【机票代理商订单】订单搜索 API请求
// taobao.jipiao.agent.order.search
//
// 卖家根据条件查询淘宝订单id列表
type TaobaoJipiaoAgentOrderSearchAPIRequest struct {
	model.Params
	// 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整）
	_beginTime string
	// 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整）
	_endTime string
	// 扩展字段:<br/>needFilterAutobook：默认true。待出票状态下，会根据此值过滤掉自动出票状态下订单，以防止重复出票的问题。对于精选票，此值需要设置成false，并由API使用者保证不会重复出票。
	_extra string
	// 订单状态，默认为空，查询所有状态的订单<br/>1:待确认<br/>2:待出票<br/>3:强制成功<br/>4:未付款<br/>5:预订成功<br/>6:已失效
	_status int64
	// 航程类型： 0.单程；1.往返
	_tripType int64
	// 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext
	_page int64
	// 是否需要行程单，true表示需要行程单；false表示不许要
	_hasItinerary bool
}

// NewTaobaoJipiaoAgentOrderSearchRequest 初始化TaobaoJipiaoAgentOrderSearchAPIRequest对象
func NewTaobaoJipiaoAgentOrderSearchRequest() *TaobaoJipiaoAgentOrderSearchAPIRequest {
	return &TaobaoJipiaoAgentOrderSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetApiMethodName() string {
	return "taobao.jipiao.agent.order.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBeginTime is BeginTime Setter
// 创建订单时间范围的开始时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认开始时间为当前时间往前推三天 （具体天数可能调整）
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 创建订单时间范围的结束时间，注意：当前搜索条件开始结束时间范围不能超过三天，默认为当前时间 （具体天数可能调整）
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetExtra is Extra Setter
// 扩展字段:&lt;br/&gt;needFilterAutobook：默认true。待出票状态下，会根据此值过滤掉自动出票状态下订单，以防止重复出票的问题。对于精选票，此值需要设置成false，并由API使用者保证不会重复出票。
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetExtra() string {
	return r._extra
}

// SetStatus is Status Setter
// 订单状态，默认为空，查询所有状态的订单&lt;br/&gt;1:待确认&lt;br/&gt;2:待出票&lt;br/&gt;3:强制成功&lt;br/&gt;4:未付款&lt;br/&gt;5:预订成功&lt;br/&gt;6:已失效
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetStatus() int64 {
	return r._status
}

// SetTripType is TripType Setter
// 航程类型： 0.单程；1.往返
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetTripType(_tripType int64) error {
	r._tripType = _tripType
	r.Set("trip_type", _tripType)
	return nil
}

// GetTripType TripType Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetTripType() int64 {
	return r._tripType
}

// SetPage is Page Setter
// 页码，默认第一页；注意：页大小固定，搜索结果中返回页大小pageSize，和是否包含下一页hasNext
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetPage() int64 {
	return r._page
}

// SetHasItinerary is HasItinerary Setter
// 是否需要行程单，true表示需要行程单；false表示不许要
func (r *TaobaoJipiaoAgentOrderSearchAPIRequest) SetHasItinerary(_hasItinerary bool) error {
	r._hasItinerary = _hasItinerary
	r.Set("has_itinerary", _hasItinerary)
	return nil
}

// GetHasItinerary HasItinerary Getter
func (r TaobaoJipiaoAgentOrderSearchAPIRequest) GetHasItinerary() bool {
	return r._hasItinerary
}
