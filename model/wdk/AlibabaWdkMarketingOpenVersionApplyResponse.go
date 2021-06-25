package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
数据同步版本号申请 APIResponse
alibaba.wdk.marketing.open.version.apply

数据同步版本号申请
*/
type AlibabaWdkMarketingOpenVersionApplyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingOpenVersionApplyResponse `json:"alibaba_wdk_marketing_open_version_apply_response,omitempty"`
}

type AlibabaWdkMarketingOpenVersionApplyResponse struct {

    // 版本号申请结果
    Result   *WdkMarketOpenResult `json:"result,omitempty"`

}
