package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensoropenwindowAPIResponse 客户端打开新页面 API返回值
// alibaba.interact.sensor.openwindow
//
// 客户端打开新页面
type AlibabainteractsensoropenwindowAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensoropenwindowAPIResponseModel
}

// AlibabainteractsensoropenwindowAPIResponseModel is 客户端打开新页面 成功返回结果
type AlibabainteractsensoropenwindowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_openwindow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
