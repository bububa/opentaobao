package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
版本数量查询 APIResponse
alibaba.wdk.marketing.open.version.count

版本数量查询
*/
type AlibabaWdkMarketingOpenVersionCountAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingOpenVersionCountResponse `json:"alibaba_wdk_marketing_open_version_count_response,omitempty"`
}

type AlibabaWdkMarketingOpenVersionCountResponse struct {

    // 查询结果
    Result   *WdkMarketOpenResult `json:"result,omitempty"`

}
