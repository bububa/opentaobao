package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系列品商品变更-移除商品 APIResponse
alibaba.wdk.series.sku.remove

系列品商品变更-移除商品
*/
type AlibabaWdkSeriesSkuRemoveAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSeriesSkuRemoveResponse `json:"alibaba_wdk_series_sku_remove_response,omitempty"`
}

type AlibabaWdkSeriesSkuRemoveResponse struct {

    // 调用结果
    ApiResult   *AlibabaWdkSeriesSkuRemoveApiResult `json:"api_result,omitempty"`

}
