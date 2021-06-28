package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
活动数据同步 APIResponse
alibaba.wdk.marketing.open.darunfa.activity.sync

大润发活动数据同步
*/
type AlibabaWdkMarketingOpenDarunfaActivitySyncAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingOpenDarunfaActivitySyncResponse
}

type AlibabaWdkMarketingOpenDarunfaActivitySyncResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_open_darunfa_activity_sync_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
