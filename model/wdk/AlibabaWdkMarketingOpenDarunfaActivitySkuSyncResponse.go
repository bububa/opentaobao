package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
营销商品数据同步 APIResponse
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkMarketingOpenDarunfaActivitySkuSyncResponse `json:"alibaba_wdk_marketing_open_darunfa_activity_sku_sync_response,omitempty"`
}

type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncResponse struct {

    // 返回结果信息
    Result   *WdkMarketOpenResult `json:"result,omitempty"`

}
