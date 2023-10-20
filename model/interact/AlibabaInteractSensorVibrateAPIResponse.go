package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorvibrateAPIResponse 客户端震动 API返回值
// alibaba.interact.sensor.vibrate
//
// 客户端震动
type AlibabainteractsensorvibrateAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensorvibrateAPIResponseModel
}

// AlibabainteractsensorvibrateAPIResponseModel is 客户端震动 成功返回结果
type AlibabainteractsensorvibrateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_vibrate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
