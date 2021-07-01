package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorNetworkstatusAPIResponse
网络状态 API返回值
alibaba.interact.sensor.networkstatus

客户端网络状态 */
type AlibabaInteractSensorNetworkstatusAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorNetworkstatusAPIResponseModel
}

// AlibabaInteractSensorNetworkstatusAPIResponseModel is 网络状态 成功返回结果
type AlibabaInteractSensorNetworkstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_networkstatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
