package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorGmediaAPIResponse
gmedia API返回值
alibaba.interact.sensor.gmedia

媒体功能 */
type AlibabaInteractSensorGmediaAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGmediaAPIResponseModel
}

// AlibabaInteractSensorGmediaAPIResponseModel is gmedia 成功返回结果
type AlibabaInteractSensorGmediaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gmedia_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
