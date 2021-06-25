package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交易组件 APIResponse
alibaba.interact.sensor.trade

交易流程
*/
type AlibabaInteractSensorTradeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorTradeResponse `json:"alibaba_interact_sensor_trade_response,omitempty"`
}

type AlibabaInteractSensorTradeResponse struct {

    // result=1
    Result   string `json:"result,omitempty"`

}
