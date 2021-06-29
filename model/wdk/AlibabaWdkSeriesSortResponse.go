package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品-商品排序 API返回值 
alibaba.wdk.series.sort

系列品商品变更-商品排序
*/
type AlibabaWdkSeriesSortAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSeriesSortResponse
}

// 系列品-商品排序 成功返回结果
type AlibabaWdkSeriesSortResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_series_sort_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    ApiResult   *AlibabaWdkSeriesSortApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
