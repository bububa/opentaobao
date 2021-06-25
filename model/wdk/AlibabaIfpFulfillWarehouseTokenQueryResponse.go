package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
同城令牌打印接口 APIResponse
alibaba.ifp.fulfill.warehouse.token.query

用于仓内作业打印包裹信息
*/
type AlibabaIfpFulfillWarehouseTokenQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaIfpFulfillWarehouseTokenQueryResponse `json:"alibaba_ifp_fulfill_warehouse_token_query_response,omitempty"`
}

type AlibabaIfpFulfillWarehouseTokenQueryResponse struct {

    // 返回结果
    WorkResult   *WorkResult `json:"work_result,omitempty"`

}
