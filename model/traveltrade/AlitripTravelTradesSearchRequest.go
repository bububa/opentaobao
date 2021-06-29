package traveltrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
飞猪度假-订单列表搜索接口 APIRequest
alitrip.travel.trades.search

订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。
*/
type AlitripTravelTradesSearchRequest struct {
    model.Params

    // 页面大小，最大支持的页面大小为100。如查询旅行购订单，则最大支持的页面大小为30
    pageSize   int64 

    // 订单创建 结束时间
    endCreatedTime   string 

    // 订单状态 过滤。1-等待买家付款，2-等待卖家发货（买家已付款），3-等待买家确认收货，4-交易关闭（买家发起的退款），6-交易成功，8-交易关闭（订单超时 自动关单）
    orderStatus   int64 

    // 当前页
    currentPage   int64 

    // 订单创建 开始时间
    startCreatedTime   string 

    // 类目筛选, 1、旅行购，旅行购定制专用字段，表示搜索旅行购订单。
    category   int64 

}

func NewAlitripTravelTradesSearchRequest() *AlitripTravelTradesSearchRequest{
    return &AlitripTravelTradesSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTravelTradesSearchRequest) GetApiMethodName() string {
    return "alitrip.travel.trades.search"
}

func (r AlitripTravelTradesSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTravelTradesSearchRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AlitripTravelTradesSearchRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AlitripTravelTradesSearchRequest) SetEndCreatedTime(endCreatedTime string) error {
    r.endCreatedTime = endCreatedTime
    r.Set("end_created_time", endCreatedTime)
    return nil
}

func (r AlitripTravelTradesSearchRequest) GetEndCreatedTime() string {
    return r.endCreatedTime
}

func (r *AlitripTravelTradesSearchRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

func (r AlitripTravelTradesSearchRequest) GetOrderStatus() int64 {
    return r.orderStatus
}

func (r *AlitripTravelTradesSearchRequest) SetCurrentPage(currentPage int64) error {
    r.currentPage = currentPage
    r.Set("current_page", currentPage)
    return nil
}

func (r AlitripTravelTradesSearchRequest) GetCurrentPage() int64 {
    return r.currentPage
}

func (r *AlitripTravelTradesSearchRequest) SetStartCreatedTime(startCreatedTime string) error {
    r.startCreatedTime = startCreatedTime
    r.Set("start_created_time", startCreatedTime)
    return nil
}

func (r AlitripTravelTradesSearchRequest) GetStartCreatedTime() string {
    return r.startCreatedTime
}

func (r *AlitripTravelTradesSearchRequest) SetCategory(category int64) error {
    r.category = category
    r.Set("category", category)
    return nil
}

func (r AlitripTravelTradesSearchRequest) GetCategory() int64 {
    return r.category
}

