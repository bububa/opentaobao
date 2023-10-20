package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorshakeAPIResponse 摇一摇 API返回值
// alibaba.interact.sensor.shake
//
// 摇一摇
type AlibabainteractsensorshakeAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensorshakeAPIResponseModel
}

// AlibabainteractsensorshakeAPIResponseModel is 摇一摇 成功返回结果
type AlibabainteractsensorshakeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_shake_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
