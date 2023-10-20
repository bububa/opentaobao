package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorSocialAPIResponse 社交组件 API返回值
// alibaba.interact.sensor.social
//
// 赞，评论 ，关注 新增接口
type AlibabaInteractSensorSocialAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorSocialAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorSocialAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorSocialAPIResponseModel).Reset()
}

// AlibabaInteractSensorSocialAPIResponseModel is 社交组件 成功返回结果
type AlibabaInteractSensorSocialAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_social_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=1
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorSocialAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorSocialAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorSocialAPIResponse)
	},
}

// GetAlibabaInteractSensorSocialAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorSocialAPIResponse
func GetAlibabaInteractSensorSocialAPIResponse() *AlibabaInteractSensorSocialAPIResponse {
	return poolAlibabaInteractSensorSocialAPIResponse.Get().(*AlibabaInteractSensorSocialAPIResponse)
}

// ReleaseAlibabaInteractSensorSocialAPIResponse 将 AlibabaInteractSensorSocialAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorSocialAPIResponse(v *AlibabaInteractSensorSocialAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorSocialAPIResponse.Put(v)
}
