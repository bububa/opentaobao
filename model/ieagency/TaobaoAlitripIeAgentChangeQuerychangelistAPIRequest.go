package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest
卖家查询改签列表 API请求
taobao.alitrip.ie.agent.change.querychangelist

提供B2B卖家查看改签列表服务 */
type TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest struct {
	model.Params
	// WAITING_CONFIRM(10, "卖家待确认"),CONFIRMED(20, "卖家已确认"),WAITING_ISSUE(30, "卖家待出票"),FROZEN_ORDER(40, "出票超时冻结"),ISSUE_SUCCESS(50, "出票成功"),CHECKING_FAILURE(60,"验真失败"),CHECKING_SUCCCESS(61,"验真成功"),REFUSED(70, "卖家已拒绝")
	_changeBizStatusEnum string
	// 改签申请单ID
	_changeOrderId int64
	// 申请原因类型 0－因乘客个人原因(自愿改签) ,1－因航班取消/延误(非自愿),
	_changeReasonType string
	// 1
	_endCreateDate string
	// 订单ID
	_orderId int64
	// 分页索引
	_pageIndex int64
	// 分页大小
	_pageSize int64
	// 排序
	_sortField int64
	// 1
	_startCreateDate string
}

// NewTaobaoAlitripIeAgentChangeQuerychangelistRequest 初始化TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest对象
func NewTaobaoAlitripIeAgentChangeQuerychangelistRequest() *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest {
	return &TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.change.querychangelist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ChangeBizStatusEnum Setter
// WAITING_CONFIRM(10, "卖家待确认"),CONFIRMED(20, "卖家已确认"),WAITING_ISSUE(30, "卖家待出票"),FROZEN_ORDER(40, "出票超时冻结"),ISSUE_SUCCESS(50, "出票成功"),CHECKING_FAILURE(60,"验真失败"),CHECKING_SUCCCESS(61,"验真成功"),REFUSED(70, "卖家已拒绝")
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetChangeBizStatusEnum(_changeBizStatusEnum string) error {
	r._changeBizStatusEnum = _changeBizStatusEnum
	r.Set("change_biz_status_enum", _changeBizStatusEnum)
	return nil
}

// Get ChangeBizStatusEnum Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetChangeBizStatusEnum() string {
	return r._changeBizStatusEnum
}

// Set is ChangeOrderId Setter
// 改签申请单ID
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetChangeOrderId(_changeOrderId int64) error {
	r._changeOrderId = _changeOrderId
	r.Set("change_order_id", _changeOrderId)
	return nil
}

// Get ChangeOrderId Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetChangeOrderId() int64 {
	return r._changeOrderId
}

// Set is ChangeReasonType Setter
// 申请原因类型 0－因乘客个人原因(自愿改签) ,1－因航班取消/延误(非自愿),
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetChangeReasonType(_changeReasonType string) error {
	r._changeReasonType = _changeReasonType
	r.Set("change_reason_type", _changeReasonType)
	return nil
}

// Get ChangeReasonType Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetChangeReasonType() string {
	return r._changeReasonType
}

// Set is EndCreateDate Setter
// 1
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetEndCreateDate(_endCreateDate string) error {
	r._endCreateDate = _endCreateDate
	r.Set("end_create_date", _endCreateDate)
	return nil
}

// Get EndCreateDate Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetEndCreateDate() string {
	return r._endCreateDate
}

// Set is OrderId Setter
// 订单ID
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is PageIndex Setter
// 分页索引
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// Set is PageSize Setter
// 分页大小
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// Set is SortField Setter
// 排序
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetSortField(_sortField int64) error {
	r._sortField = _sortField
	r.Set("sort_field", _sortField)
	return nil
}

// Get SortField Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetSortField() int64 {
	return r._sortField
}

// Set is StartCreateDate Setter
// 1
func (r *TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) SetStartCreateDate(_startCreateDate string) error {
	r._startCreateDate = _startCreateDate
	r.Set("start_create_date", _startCreateDate)
	return nil
}

// Get StartCreateDate Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistAPIRequest) GetStartCreateDate() string {
	return r._startCreateDate
}
