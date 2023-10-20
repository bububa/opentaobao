package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoweikeeserviceordergetAPIRequest 客服外包订单查询 API请求
// taobao.weike.eservice.order.get
//
// 用于客服外包中服务商查询订单列表
type TaobaoweikeeserviceordergetAPIRequest struct {
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

// NewTaobaoweikeeserviceordergetRequest 初始化TaobaoweikeeserviceordergetAPIRequest对象
func NewTaobaoweikeeserviceordergetRequest() *TaobaoweikeeserviceordergetAPIRequest {
	return &TaobaoweikeeserviceordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoweikeeserviceordergetAPIRequest) GetApiMethodName() string {
	return "taobao.weike.eservice.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoweikeeserviceordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoweikeeserviceordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderIdList is OrderIdList Setter
// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
func (r *TaobaoweikeeserviceordergetAPIRequest) SetOrderIdList(_orderIdList []int64) error {
	r._orderIdList = _orderIdList
	r.Set("order_id_list", _orderIdList)
	return nil
}

// GetOrderIdList OrderIdList Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetOrderIdList() []int64 {
	return r._orderIdList
}

// SetStartDate is StartDate Setter
// 订单服务开始日期
func (r *TaobaoweikeeserviceordergetAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetSellerNick is SellerNick Setter
// 商家昵称
func (r *TaobaoweikeeserviceordergetAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetEndDate is EndDate Setter
// 订单服务结束日期
func (r *TaobaoweikeeserviceordergetAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageSize is PageSize Setter
// 每页记录数（默认20，最大不超过20）
func (r *TaobaoweikeeserviceordergetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNum is PageNum Setter
// 页码（默认为1）
func (r *TaobaoweikeeserviceordergetAPIRequest) SetPageNum(_pageNum int64) error {
	r._pageNum = _pageNum
	r.Set("page_num", _pageNum)
	return nil
}

// GetPageNum PageNum Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetPageNum() int64 {
	return r._pageNum
}

// SetSchedulingState is SchedulingState Setter
// 订单是否可以排班
func (r *TaobaoweikeeserviceordergetAPIRequest) SetSchedulingState(_schedulingState bool) error {
	r._schedulingState = _schedulingState
	r.Set("scheduling_state", _schedulingState)
	return nil
}

// GetSchedulingState SchedulingState Getter
func (r TaobaoweikeeserviceordergetAPIRequest) GetSchedulingState() bool {
	return r._schedulingState
}
