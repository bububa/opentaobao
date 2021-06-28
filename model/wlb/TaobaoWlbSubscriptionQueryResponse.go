package wlb

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询商家定购的所有服务 APIResponse
taobao.wlb.subscription.query

查询商家定购的所有服务,可通过入参状态来筛选
*/
type TaobaoWlbSubscriptionQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoWlbSubscriptionQueryResponse `json:"wlb_subscription_query_response,omitempty"` 
    TaobaoWlbSubscriptionQueryResponse
}

/* model for simplify = false
type TaobaoWlbSubscriptionQueryResponse struct {

    // 结果总数
    
    TotalCount   int64 `json:"total_count,omitempty"`
    

    // 卖家定购的服务列表
    
    SellerSubscriptionList  struct {
        WlbSellerSubscription  []WlbSellerSubscription `json:"wlb_seller_subscription,omitempty"`
    } `json:"seller_subscription_list,omitempty"`
    

}
*/

type TaobaoWlbSubscriptionQueryResponse struct {

    // 结果总数
    TotalCount   int64 `json:"total_count,omitempty"`

    // 卖家定购的服务列表
    SellerSubscriptionList   []WlbSellerSubscription `json:"seller_subscription_list,omitempty"`

}
