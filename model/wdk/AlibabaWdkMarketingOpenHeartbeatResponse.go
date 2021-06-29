package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
心跳服务【10s一次】 API返回值 
alibaba.wdk.marketing.open.heartbeat

商家数据同步心跳服务
*/
type AlibabaWdkMarketingOpenHeartbeatAPIResponse struct {
    model.CommonResponse
    AlibabaWdkMarketingOpenHeartbeatResponse
}

// 心跳服务【10s一次】 成功返回结果
type AlibabaWdkMarketingOpenHeartbeatResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_marketing_open_heartbeat_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果信息
    Result   *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
