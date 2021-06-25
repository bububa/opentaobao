package cityretail

import (
    "github.com/bububa/opentaobao/model"
)

/* 
完美履约订单物流状态查询接口 APIResponse
tmall.cityretail.wmfl.order.logistics.query

完美履约订单物流状态查询接口，该接口只能查询未完结的履约单以及完结的3天内订单
*/
type TmallCityretailWmflOrderLogisticsQueryAPIResponse struct {
    model.CommonResponse
    Response *TmallCityretailWmflOrderLogisticsQueryResponse `json:"tmall_cityretail_wmfl_order_logistics_query_response,omitempty"`
}

type TmallCityretailWmflOrderLogisticsQueryResponse struct {

    // 返回对象
    Result   *WorkResult `json:"result,omitempty"`

}
