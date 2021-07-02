package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceOrderGetAPIRequest 客服外包订单查询 API请求
// taobao.weike.eservice.order.get
//
// 用于客服外包中服务商查询订单列表
type TaobaoWeikeEserviceOrderGetAPIRequest struct {
	model.Params
	// 订单服务开始日期
	_startDate string
	// 订单是否可以排班
	_schedulingState bool
	// 商家昵称
	_sellerNick string
	// 每页记录数（默认20，最大不超过20）
	_pageSize int64
	// 订单服务结束日期
	_endDate string
	// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
	_orderIdList []int64
	// 页码（默认为1）
	_pageNum int64
}

// NewTaobaoWeikeEserviceOrderGetRequest 初始化TaobaoWeikeEserviceOrderGetAPIRequest对象
func NewTaobaoWeikeEserviceOrderGetRequest() *TaobaoWeikeEserviceOrderGetAPIRequest {
	return &TaobaoWeikeEserviceOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is StartDate Setter
// 订单服务开始日期
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// Set is SchedulingState Setter
// 订单是否可以排班
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetSchedulingState(_schedulingState bool) error {
	r._schedulingState = _schedulingState
	r.Set("scheduling_state", _schedulingState)
	return nil
}

// Get SchedulingState Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetSchedulingState() bool {
	return r._schedulingState
}

// Set is SellerNick Setter
// 商家昵称
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// Get SellerNick Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// Set is PageSize Setter
// 每页记录数（默认20，最大不超过20）
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is EndDate Setter
// 订单服务结束日期
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// Set is OrderIdList Setter
// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetOrderIdList(_orderIdList []int64) error {
	r._orderIdList = _orderIdList
	r.Set("order_id_list", _orderIdList)
	return nil
}

// Get OrderIdList Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetOrderIdList() []int64 {
	return r._orderIdList
}

// Set is PageNum Setter
// 页码（默认为1）
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// Get PageNum Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}
