package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
活动数据同步 APIResponse
alibaba.wdk.marketing.open.darunfa.activity.sync

大润发活动数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingOpenDarunfaActivitySyncResponse `json:"alibaba_wdk_marketing_open_darunfa_activity_sync_response,omitempty"` 
    AlibabaWdkMarketingOpenDarunfaActivitySyncResponse
}

/* model for simplify = false
type AlibabaWdkMarketingOpenDarunfaActivitySyncResponse struct {

    // 结果信息
    
    Result  *struct {
        WdkMarketOpenResult  *WdkMarketOpenResult `json:"wdk_market_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingOpenDarunfaActivitySyncResponse struct {

    // 结果信息
    Result   *WdkMarketOpenResult `json:"result,omitempty"`

}
