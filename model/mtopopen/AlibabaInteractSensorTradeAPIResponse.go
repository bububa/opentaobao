package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorTradeAPIResponse
交易组件 API返回值
alibaba.interact.sensor.trade

交易流程 */
type AlibabaInteractSensorTradeAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorTradeAPIResponseModel
}

// AlibabaInteractSensorTradeAPIResponseModel is 交易组件 成功返回结果
type AlibabaInteractSensorTradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=1
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
