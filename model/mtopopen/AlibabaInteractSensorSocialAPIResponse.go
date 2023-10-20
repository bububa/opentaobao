package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractsensorsocialAPIResponse 社交组件 API返回值
// alibaba.interact.sensor.social
//
// 赞，评论 ，关注 新增接口
type AlibabainteractsensorsocialAPIResponse struct {
	model.CommonResponse
	AlibabainteractsensorsocialAPIResponseModel
}

// AlibabainteractsensorsocialAPIResponseModel is 社交组件 成功返回结果
type AlibabainteractsensorsocialAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_social_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=1
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
