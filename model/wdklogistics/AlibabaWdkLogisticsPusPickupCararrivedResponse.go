package wdklogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自提业务-车辆到达上报车牌号 APIResponse
alibaba.wdk.logistics.pus.pickup.cararrived

自提业务-汽车自提,车辆到达上报车牌号
*/
type AlibabaWdkLogisticsPusPickupCararrivedAPIResponse struct {
    model.CommonResponse
    AlibabaWdkLogisticsPusPickupCararrivedResponse
}

type AlibabaWdkLogisticsPusPickupCararrivedResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_logistics_pus_pickup_cararrived_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result根结点
    
    Result   *LogisticsResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
