package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-移除商品 APIResponse
alibaba.wdk.series.sku.remove

系列品商品变更-移除商品
*/
type AlibabaWdkSeriesSkuRemoveAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSeriesSkuRemoveResponse
}

type AlibabaWdkSeriesSkuRemoveResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_series_sku_remove_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 调用结果
    
    ApiResult   *AlibabaWdkSeriesSkuRemoveApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`

    
}
