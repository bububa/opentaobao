package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
心跳服务【10s一次】 APIResponse
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务
*/
type AlibabaWdkMarketingOpenHeartbeatAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_marketing_open_heartbeat_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果信息
    
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"