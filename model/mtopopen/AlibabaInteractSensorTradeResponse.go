package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
交易组件 APIResponse
alibaba.interact.sensor.trade

交易流程
*/
type AlibabaInteractSensorTradeAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorTradeResponse
}

type AlibabaInteractSensorTradeResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_trade_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result=1
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
