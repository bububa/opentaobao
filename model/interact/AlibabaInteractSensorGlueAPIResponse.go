package interact

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorglueAPIResponse 视频播放 API返回值
// alibaba.interact.sensor.glue
//
// 视频播放
type AlibabainteractsensorglueAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensorglueAPIResponseModel
}

// AlibabainteractsensorglueAPIResponseModel is 视频播放 成功返回结果
type AlibabainteractsensorglueAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_glue_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
