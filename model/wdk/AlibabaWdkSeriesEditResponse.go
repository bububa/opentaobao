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
    AlibabaWdkSeriesEditResponse
}

type AlibabaWdkSeriesEditResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_series_edit_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    ApiResult   *AlibabaWdkSeriesEditApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
