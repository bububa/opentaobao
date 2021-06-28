package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘下单能力开放 APIResponse
alibaba.interact.sensor.trade.buy

交易流程鉴权
*/
type AlibabaInteractSensorTradeBuyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_interact_sensor_trade_buy_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

}
