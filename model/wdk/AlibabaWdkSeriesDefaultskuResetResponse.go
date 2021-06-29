package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-重置默认商品 APIResponse
alibaba.wdk.series.defaultsku.reset

系列品商品变更-重置默认商品
*/
type AlibabaWdkSeriesDefaultskuResetAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSeriesDefaultskuResetResponse
}

type AlibabaWdkSeriesDefaultskuResetResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_series_defaultsku_reset_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果
    
    ApiResult   *AlibabaWdkSeriesDefaultskuResetApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
