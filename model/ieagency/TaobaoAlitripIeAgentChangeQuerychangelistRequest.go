package ieagency

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家查询改签列表 API请求
taobao.alitrip.ie.agent.change.querychangelist

提供B2B卖家查看改签列表服务
*/
type TaobaoAlitripIeAgentChangeQuerychangelistRequest struct {
    model.Params
    // WAITING_CONFIRM(10, "卖家待确认"),CONFIRMED(20, "卖家已确认"),WAITING_ISSUE(30, "卖家待出票"),FROZEN_ORDER(40, "出票超时冻结"),ISSUE_SUCCESS(50, "出票成功"),CHECKING_FAILURE(60,"验真失败"),CHECKING_SUCCCESS(61,"验真成功"),REFUSED(70, "卖家已拒绝")
    changeBizStatusEnum   string
    // 改签申请单ID
    changeOrderId   int64
    // 申请原因类型 0－因乘客个人原因(自愿改签) ,1－因航班取消/延误(非自愿),
    changeReasonType   string
    // 1
    endCreateDate   string
    // 订单ID
    orderId   int64
    // 分页索引
    pageIndex   int64
    // 分页大小
    pageSize   int64
    // 排序
    sortField   int64
    // 1
    startCreateDate   string
}

// 初始化TaobaoAlitripIeAgentChangeQuerychangelistRequest对象
func NewTaobaoAlitripIeAgentChangeQuerychangelistRequest() *TaobaoAlitripIeAgentChangeQuerychangelistRequest{
    return &TaobaoAlitripIeAgentChangeQuerychangelistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetApiMethodName() string {
    return "taobao.alitrip.ie.agent.change.querychangelist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChangeBizStatusEnum Setter
// WAITING_CONFIRM(10, "卖家待确认"),CONFIRMED(20, "卖家已确认"),WAITING_ISSUE(30, "卖家待出票"),FROZEN_ORDER(40, "出票超时冻结"),ISSUE_SUCCESS(50, "出票成功"),CHECKING_FAILURE(60,"验真失败"),CHECKING_SUCCCESS(61,"验真成功"),REFUSED(70, "卖家已拒绝")
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetChangeBizStatusEnum(changeBizStatusEnum string) error {
    r.changeBizStatusEnum = changeBizStatusEnum
    r.Set("change_biz_status_enum", changeBizStatusEnum)
    return nil
}

// ChangeBizStatusEnum Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetChangeBizStatusEnum() string {
    return r.changeBizStatusEnum
}
// ChangeOrderId Setter
// 改签申请单ID
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetChangeOrderId(changeOrderId int64) error {
    r.changeOrderId = changeOrderId
    r.Set("change_order_id", changeOrderId)
    return nil
}

// ChangeOrderId Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetChangeOrderId() int64 {
    return r.changeOrderId
}
// ChangeReasonType Setter
// 申请原因类型 0－因乘客个人原因(自愿改签) ,1－因航班取消/延误(非自愿),
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetChangeReasonType(changeReasonType string) error {
    r.changeReasonType = changeReasonType
    r.Set("change_reason_type", changeReasonType)
    return nil
}

// ChangeReasonType Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetChangeReasonType() string {
    return r.changeReasonType
}
// EndCreateDate Setter
// 1
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetEndCreateDate(endCreateDate string) error {
    r.endCreateDate = endCreateDate
    r.Set("end_create_date", endCreateDate)
    return nil
}

// EndCreateDate Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetEndCreateDate() string {
    return r.endCreateDate
}
// OrderId Setter
// 订单ID
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetOrderId() int64 {
    return r.orderId
}
// PageIndex Setter
// 分页索引
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetPageIndex(pageIndex int64) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetPageIndex() int64 {
    return r.pageIndex
}
// PageSize Setter
// 分页大小
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetPageSize() int64 {
    return r.pageSize
}
// SortField Setter
// 排序
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetSortField(sortField int64) error {
    r.sortField = sortField
    r.Set("sort_field", sortField)
    return nil
}

// SortField Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetSortField() int64 {
    return r.sortField
}
// StartCreateDate Setter
// 1
func (r *TaobaoAlitripIeAgentChangeQuerychangelistRequest) SetStartCreateDate(startCreateDate string) error {
    r.startCreateDate = startCreateDate
    r.Set("start_create_date", startCreateDate)
    return nil
}

// StartCreateDate Getter
func (r TaobaoAlitripIeAgentChangeQuerychangelistRequest) GetStartCreateDate() string {
    return r.startCreateDate
}
