package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-添加商品 APIResponse
alibaba.wdk.series.sku.add

系列品商品变更-添加商品
*/
type AlibabaWdkSeriesSkuAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_series_sku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    ApiResult   *AlibabaWdkSeriesSkuAddApiResult `json:"api_result,omitempty" xml:"