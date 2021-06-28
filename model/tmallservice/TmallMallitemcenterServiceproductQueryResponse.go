package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务商服务产品查询 APIResponse
tmall.mallitemcenter.serviceproduct.query

查询天猫服务的服务商发布的服务产品
*/
type TmallMallitemcenterServiceproductQueryAPIResponse struct {
    model.CommonResponse
    // Response *TmallMallitemcenterServiceproductQueryResponse `json:"tmall_mallitemcenter_serviceproduct_query_response,omitempty"` 
    TmallMallitemcenterServiceproductQueryResponse
}

/* model for simplify = false
type TmallMallitemcenterServiceproductQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        TmallMallitemcenterServiceproductQueryResult  *TmallMallitemcenterServiceproductQueryResult `json:"tmall_mallitemcenter_serviceproduct_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TmallMallitemcenterServiceproductQueryResponse struct {

    // 接口返回model
    Result   *TmallMallitemcenterServiceproductQueryResult `json:"result,omitempty"`

}
