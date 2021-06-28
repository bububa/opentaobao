package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
心跳服务【10s一次】 APIResponse
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务
*/
type AlibabaWdkMarketingOpenHeartbeatAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWdkMarketingOpenHeartbeatResponse `json:"alibaba_wdk_marketing_open_heartbeat_response,omitempty"` 
    AlibabaWdkMarketingOpenHeartbeatResponse
}

/* model for simplify = false
type AlibabaWdkMarketingOpenHeartbeatResponse struct {

    // 结果信息
    
    Result  *struct {
        WdkMarketOpenResult  *WdkMarketOpenResult `json:"wdk_market_open_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaWdkMarketingOpenHeartbeatResponse struct {

    // 结果信息
    Result   *WdkMarketOpenResult `json:"result,omitempty"`

}
