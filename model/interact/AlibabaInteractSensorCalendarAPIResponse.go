package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorCalendarAPIResponse 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) API返回值
// alibaba.interact.sensor.calendar
//
// 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒)
type AlibabaInteractSensorCalendarAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorCalendarAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorCalendarAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorCalendarAPIResponseModel).Reset()
}

// AlibabaInteractSensorCalendarAPIResponseModel is 天猫互动游戏开放平台需要授权的传感器类接口(日历提醒) 成功返回结果
type AlibabaInteractSensorCalendarAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_calendar_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorCalendarAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Succ = false
}

var poolAlibabaInteractSensorCalendarAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorCalendarAPIResponse)
	},
}

// GetAlibabaInteractSensorCalendarAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorCalendarAPIResponse
func GetAlibabaInteractSensorCalendarAPIResponse() *AlibabaInteractSensorCalendarAPIResponse {
	return poolAlibabaInteractSensorCalendarAPIResponse.Get().(*AlibabaInteractSensorCalendarAPIResponse)
}

// ReleaseAlibabaInteractSensorCalendarAPIResponse 将 AlibabaInteractSensorCalendarAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorCalendarAPIResponse(v *AlibabaInteractSensorCalendarAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorCalendarAPIResponse.Put(v)
}
