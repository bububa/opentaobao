package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销渠道订单详情查询 APIResponse
taobao.xhotel.distribution.order.detail.search

该接口用于提供酒店分销渠道的订单详情查询
*/
type TaobaoXhotelDistributionOrderDetailSearchAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoXhotelDistributionOrderDetailSearchResponse `json:"xhotel_distribution_order_detail_search_response,omitempty"` 
    TaobaoXhotelDistributionOrderDetailSearchResponse
}

/* model for simplify = false
type TaobaoXhotelDistributionOrderDetailSearchResponse struct {

    // 错误码
    
    Error   string `json:"error,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 订单详情对象
    
    TopDistributionOrderDetail  *struct {
        TopDistributionOrderDetail  *TopDistributionOrderDetail `json:"top_distribution_order_detail,omitempty"`
    } `json:"top_distribution_order_detail,omitempty"`
    

}
*/

type TaobaoXhotelDistributionOrderDetailSearchResponse struct {

    // 错误码
    Error   string `json:"error,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 订单详情对象
    TopDistributionOrderDetail   *TopDistributionOrderDetail `json:"top_distribution_order_detail,omitempty"`

}
