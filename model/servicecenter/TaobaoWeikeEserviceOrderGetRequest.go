package servicecenter

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客服外包订单查询 API请求
taobao.weike.eservice.order.get

用于客服外包中服务商查询订单列表
*/
type TaobaoWeikeEserviceOrderGetRequest struct {
    model.Params
    // 订单服务开始日期
    startDate   string
    // 订单是否可以排班
    schedulingState   bool
    // 商家昵称
    sellerNick   string
    // 每页记录数（默认20，最大不超过20）
    pageSize   int64
    // 订单服务结束日期
    endDate   string
    // 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
    orderIdList   []int64
    // 页码（默认为1）
    pageNum   int64
}

// 初始化TaobaoWeikeEserviceOrderGetRequest对象
func NewTaobaoWeikeEserviceOrderGetRequest() *TaobaoWeikeEserviceOrderGetRequest{
    return &TaobaoWeikeEserviceOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWeikeEserviceOrderGetRequest) GetApiMethodName() string {
    return "taobao.weike.eservice.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWeikeEserviceOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 订单服务开始日期
func (r *TaobaoWeikeEserviceOrderGetRequest) SetStartDate(startDate string) error {
    r.startDate = startDate
    r.Set("start_date", startDate)
    return nil
}

// StartDate Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetStartDate() string {
    return r.startDate
}
// SchedulingState Setter
// 订单是否可以排班
func (r *TaobaoWeikeEserviceOrderGetRequest) SetSchedulingState(schedulingState bool) error {
    r.schedulingState = schedulingState
    r.Set("scheduling_state", schedulingState)
    return nil
}

// SchedulingState Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetSchedulingState() bool {
    return r.schedulingState
}
// SellerNick Setter
// 商家昵称
func (r *TaobaoWeikeEserviceOrderGetRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

// SellerNick Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetSellerNick() string {
    return r.sellerNick
}
// PageSize Setter
// 每页记录数（默认20，最大不超过20）
func (r *TaobaoWeikeEserviceOrderGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// EndDate Setter
// 订单服务结束日期
func (r *TaobaoWeikeEserviceOrderGetRequest) SetEndDate(endDate string) error {
    r.endDate = endDate
    r.Set("end_date", endDate)
    return nil
}

// EndDate Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetEndDate() string {
    return r.endDate
}
// OrderIdList Setter
// 订单ID列表，最大不超过20个（这个参数指定后，其它过滤条件失效）
func (r *TaobaoWeikeEserviceOrderGetRequest) SetOrderIdList(orderIdList []int64) error {
    r.orderIdList = orderIdList
    r.Set("order_id_list", orderIdList)
    return nil
}

// OrderIdList Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetOrderIdList() []int64 {
    return r.orderIdList
}
// PageNum Setter
// 页码（默认为1）
func (r *TaobaoWeikeEserviceOrderGetRequest) SetPageNum(pageNum int64) error {
    r.pageNum = pageNum
    r.Set("page_num", pageNum)
    return nil
}

// PageNum Getter
func (r TaobaoWeikeEserviceOrderGetRequest) GetPageNum() int64 {
    return r.pageNum
}
