package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensortoastAPIResponse toast API返回值
// alibaba.interact.sensor.toast
//
// toast提示
type AlibabainteractsensortoastAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensortoastAPIResponseModel
}

// AlibabainteractsensortoastAPIResponseModel is toast 成功返回结果
type AlibabainteractsensortoastAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_toast_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
