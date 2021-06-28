package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-更新系列 APIResponse
alibaba.wdk.series.edit

系列品变更-更新系列
*/
type AlibabaWdkSeriesEditAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkSeriesEditResponse `json:"alibaba_wdk_series_edit_response,omitempty"` 
    AlibabaWdkSeriesEditResponse
}

/* model for simplify = false
type AlibabaWdkSeriesEditResponse struct {

    // 调用结果
    
    ApiResult  *struct {
        AlibabaWdkSeriesEditApiResult  *AlibabaWdkSeriesEditApiResult `json:"alibaba_wdk_series_edit_api_result,omitempty"`
    } `json:"api_result,omitempty"`
    

}
*/

type AlibabaWdkSeriesEditResponse struct {

    // 调用结果
    ApiResult   *AlibabaWdkSeriesEditApiResult `json:"api_result,omitempty"`

}
