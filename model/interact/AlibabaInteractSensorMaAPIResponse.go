package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensormaAPIResponse 码相关API API返回值
// alibaba.interact.sensor.ma
//
// 码相关API
type AlibabainteractsensormaAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensormaAPIResponseModel
}

// AlibabainteractsensormaAPIResponseModel is 码相关API 成功返回结果
type AlibabainteractsensormaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_ma_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
