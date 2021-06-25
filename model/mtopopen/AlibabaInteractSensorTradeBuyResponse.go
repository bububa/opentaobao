package mtopopen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
手淘下单能力开放 APIResponse
alibaba.interact.sensor.trade.buy

交易流程鉴权
*/
type AlibabaInteractSensorTradeBuyAPIResponse struct {
    model.CommonResponse
    Response *AlibabaInteractSensorTradeBuyResponse `json:"alibaba_interact_sensor_trade_buy_response,omitempty"`
}

type AlibabaInteractSensorTradeBuyResponse struct {

}
