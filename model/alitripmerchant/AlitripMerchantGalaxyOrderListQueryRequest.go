package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单列表查询 API请求
alitrip.merchant.galaxy.order.list.query

为C端用户提供酒店预订订单列表查询服务，包括订单支付状态、订单日期
*/
type AlitripMerchantGalaxyOrderListQueryRequest struct {
    model.Params
    // 租户标识
    tenantKey   string
    // 用户token
    token   string
    // 订单状态
    orderStatus   string
    // 入住时间
    startTime   string
    // 入住时间
    endTime   string
    // 页数
    page   int64
    // 每页行数
    row   int64
}

// 初始化AlitripMerchantGalaxyOrderListQueryRequest对象
func NewAlitripMerchantGalaxyOrderListQueryRequest() *AlitripMerchantGalaxyOrderListQueryRequest{
    return &AlitripMerchantGalaxyOrderListQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.list.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TenantKey Setter
// 租户标识
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

// TenantKey Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetTenantKey() string {
    return r.tenantKey
}
// Token Setter
// 用户token
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetToken() string {
    return r.token
}
// OrderStatus Setter
// 订单状态
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetOrderStatus(orderStatus string) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

// OrderStatus Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetOrderStatus() string {
    return r.orderStatus
}
// StartTime Setter
// 入住时间
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

// StartTime Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetStartTime() string {
    return r.startTime
}
// EndTime Setter
// 入住时间
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

// EndTime Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetEndTime() string {
    return r.endTime
}
// Page Setter
// 页数
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

// Page Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetPage() int64 {
    return r.page
}
// Row Setter
// 每页行数
func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetRow(row int64) error {
    r.row = row
    r.Set("row", row)
    return nil
}

// Row Getter
func (r AlitripMerchantGalaxyOrderListQueryRequest) GetRow() int64 {
    return r.row
}
