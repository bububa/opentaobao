package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
物流单信息查询 APIResponse
tmall.servicecenter.workcard.logisticsorder.query

物流订单信息查询API
*/
type TmallServicecenterWorkcardLogisticsorderQueryAPIResponse struct {
    model.CommonResponse
    // Response *TmallServicecenterWorkcardLogisticsorderQueryResponse `json:"tmall_servicecenter_workcard_logisticsorder_query_response,omitempty"` 
    TmallServicecenterWorkcardLogisticsorderQueryResponse
}

/* model for simplify = false
type TmallServicecenterWorkcardLogisticsorderQueryResponse struct {

    // 结果
    
    Result  *struct {
        FulfilplatformResult  *FulfilplatformResult `json:"fulfilplatform_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallServicecenterWorkcardLogisticsorderQueryResponse struct {

    // 结果
    Result   *FulfilplatformResult `json:"result,omitempty"`

}
