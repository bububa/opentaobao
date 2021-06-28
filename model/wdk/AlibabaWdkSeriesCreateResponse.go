package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品变更-新增系列 APIResponse
alibaba.wdk.series.create

系列品变更-新增系列
*/
type AlibabaWdkSeriesCreateAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSeriesCreateResponse
}

type AlibabaWdkSeriesCreateResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_series_create_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    ApiResult   *AlibabaWdkSeriesCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
