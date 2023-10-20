package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorBlowAPIResponse 吹气 API返回值
// alibaba.interact.sensor.blow
//
// 客户端吹气
type AlibabaInteractSensorBlowAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorBlowAPIResponseModel
}

// AlibabaInteractSensorBlowAPIResponseModel is 吹气 成功返回结果
type AlibabaInteractSensorBlowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_blow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0 表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
