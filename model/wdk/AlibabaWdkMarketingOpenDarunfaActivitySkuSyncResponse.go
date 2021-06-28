package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
营销商品数据同步 APIResponse
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_open_darunfa_activity_sku_sync_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果信息
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"