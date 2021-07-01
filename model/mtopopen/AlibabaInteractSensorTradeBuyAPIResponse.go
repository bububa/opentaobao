package mtopopen

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
手淘下单能力开放 API返回值 
alibaba.interact.sensor.trade.buy

交易流程鉴权
*/
type AlibabaInteractSensorTradeBuyAPIResponse struct {
    model.CommonResponse
    AlibabaInteractSensorTradeBuyAPIResponseModel
}

// 手淘下单能力开放 成功返回结果
type AlibabaInteractSensorTradeBuyAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_interact_sensor_trade_buy_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
