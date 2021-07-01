package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractSensorTitlebarhideAPIResponse
隐藏titleBar API返回值
alibaba.interact.sensor.titlebarhide

隐藏titleBar */
type AlibabaInteractSensorTitlebarhideAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorTitlebarhideAPIResponseModel
}

// AlibabaInteractSensorTitlebarhideAPIResponseModel is 隐藏titleBar 成功返回结果
type AlibabaInteractSensorTitlebarhideAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_titlebarhide_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
