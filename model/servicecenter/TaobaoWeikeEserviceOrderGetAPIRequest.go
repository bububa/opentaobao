package servicecenter

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWeikeEserviceOrderGetAPIRequest 客服外包订单查询 API请求
// taobao.weike.eservice.order.get
//
// 用于客服外包中服务商查询订单列表
type TaobaoWeikeEserviceOrderGetAPIRequest struct {
	model.Params
	// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
	_orderIdList []int64
	// 订单服务开始日期
	_startDate string
	// 商家昵称
	_sellerNick string
	// 订单服务结束日期
	_endDate string
	// 每页记录数（默认20，最大不超过20）
	_pageSize int64
	// 页码（默认为1）
	_pageNum int64
	// 订单是否可以排班
	_schedulingState bool
}

// NewTaobaoWeikeEserviceOrderGetRequest 初始化TaobaoWeikeEserviceOrderGetAPIRequest对象
func NewTaobaoWeikeEserviceOrderGetRequest() *TaobaoWeikeEserviceOrderGetAPIRequest {
	return &TaobaoWeikeEserviceOrderGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) Reset() {
	r._orderIdList = r._orderIdList[:0]
	r._startDate = ""
	r._sellerNick = ""
	r._endDate = ""
	r._pageSize = 0
	r._pageNum = 0
	r._schedulingState = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIdList is OrderIdList Setter
// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetOrderIdList(_orderIdList []int64) error {
	r._orderIdList = _orderIdList
	r.Set("order_id_list", _orderIdList)
	return nil
}

// GetOrderIdList OrderIdList Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetOrderIdList() []int64 {
	return r._orderIdList
}

// SetStartDate is StartDate Setter
// 订单服务开始日期
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetSellerNick is SellerNick Setter
// 商家昵称
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetEndDate is EndDate Setter
// 订单服务结束日期
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大不超过20）
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码（默认为1）
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetSchedulingState is SchedulingState Setter
// 订单是否可以排班
func (r *TaobaoWeikeEserviceOrderGetAPIRequest) SetSchedulingState(_schedulingState bool) error {
	r._schedulingState = _schedulingState
	r.Set("scheduling_state", _schedulingState)
	return nil
}

// GetSchedulingState SchedulingState Getter
func (r TaobaoWeikeEserviceOrderGetAPIRequest) GetSchedulingState() bool {
	return r._schedulingState
}

var poolTaobaoWeikeEserviceOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWeikeEserviceOrderGetRequest()
	},
}

// GetTaobaoWeikeEserviceOrderGetRequest 从 sync.Pool 获取 TaobaoWeikeEserviceOrderGetAPIRequest
func GetTaobaoWeikeEserviceOrderGetAPIRequest() *TaobaoWeikeEserviceOrderGetAPIRequest {
	return poolTaobaoWeikeEserviceOrderGetAPIRequest.Get().(*TaobaoWeikeEserviceOrderGetAPIRequest)
}

// ReleaseTaobaoWeikeEserviceOrderGetAPIRequest 将 TaobaoWeikeEserviceOrderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWeikeEserviceOrderGetAPIRequest(v *TaobaoWeikeEserviceOrderGetAPIRequest) {
	v.Reset()
	poolTaobaoWeikeEserviceOrderGetAPIRequest.Put(v)
}
