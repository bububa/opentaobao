package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-更新系列 APIResponse
alibaba.wdk.series.edit

系列品变更-更新系列
*/
type AlibabaWdkSeriesEditAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_series_edit_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    ApiResult   *AlibabaWdkSeriesEditApiResult `json:"api_result,omitempty" xml:"