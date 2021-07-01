package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorPopwindowAPIResponse
popwindow API返回值
alibaba.interact.sensor.popwindow

popwindow */
type AlibabaInteractSensorPopwindowAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorPopwindowAPIResponseModel
}

// AlibabaInteractSensorPopwindowAPIResponseModel is popwindow 成功返回结果
type AlibabaInteractSensorPopwindowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_popwindow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
