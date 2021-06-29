package alitripmerchant

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
星河-订单列表查询 APIRequest
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

func NewAlitripMerchantGalaxyOrderListQueryRequest() *AlitripMerchantGalaxyOrderListQueryRequest{
    return &AlitripMerchantGalaxyOrderListQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetApiMethodName() string {
    return "alitrip.merchant.galaxy.order.list.query"
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetTenantKey(tenantKey string) error {
    r.tenantKey = tenantKey
    r.Set("tenant_key", tenantKey)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetTenantKey() string {
    return r.tenantKey
}

func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetToken() string {
    return r.token
}

func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetOrderStatus(orderStatus string) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetOrderStatus() string {
    return r.orderStatus
}

func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetStartTime(startTime string) error {
    r.startTime = startTime
    r.Set("start_time", startTime)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetStartTime() string {
    return r.startTime
}

func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetEndTime(endTime string) error {
    r.endTime = endTime
    r.Set("end_time", endTime)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetEndTime() string {
    return r.endTime
}

func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetPage(page int64) error {
    r.page = page
    r.Set("page", page)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetPage() int64 {
    return r.page
}

func (r *AlitripMerchantGalaxyOrderListQueryRequest) SetRow(row int64) error {
    r.row = row
    r.Set("row", row)
    return nil
}

func (r AlitripMerchantGalaxyOrderListQueryRequest) GetRow() int64 {
    return r.row
}

