package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-重置默认商品 APIResponse
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品
*/
type AlibabaWdkSeriesDefaultskuResetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSeriesDefaultskuResetResponse `json:"alibaba_wdk_series_defaultsku_reset_response,omitempty"` 
    AlibabaWdkSeriesDefaultskuResetResponse
}

/* model for simplify = false
type AlibabaWdkSeriesDefaultskuResetResponse struct {

    // 调用结果
    
    ApiResult  *struct {
        AlibabaWdkSeriesDefaultskuResetApiResult  *AlibabaWdkSeriesDefaultskuResetApiResult `json:"alibaba_wdk_series_defaultsku_reset_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkSeriesDefaultskuResetResponse struct {

    // 调用结果
    ApiResult   *AlibabaWdkSeriesDefaultskuResetApiResult `json:"api_result,omitempty"`

}
