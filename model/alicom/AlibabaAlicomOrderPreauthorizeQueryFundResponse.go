package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
资金流水查询 APIResponse
alibaba.alicom.order.preauthorize.query.fund

预授权-资金流水查询
*/
type AlibabaAlicomOrderPreauthorizeQueryFundAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlicomOrderPreauthorizeQueryFundResponse `json:"alibaba_alicom_order_preauthorize_query_fund_response,omitempty"` 
    AlibabaAlicomOrderPreauthorizeQueryFundResponse
}

/* model for simplify = false
type AlibabaAlicomOrderPreauthorizeQueryFundResponse struct {

    // 接口结果
    
    Result  *struct {
        CommonResult  *CommonResult `json:"common_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlicomOrderPreauthorizeQueryFundResponse struct {

    // 接口结果
    Result   *CommonResult `json:"result,omitempty"`

}
