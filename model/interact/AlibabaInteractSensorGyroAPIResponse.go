package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorGyroAPIResponse 陀螺仪 API返回值
// alibaba.interact.sensor.gyro
//
// 客户端陀螺仪
type AlibabaInteractSensorGyroAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorGyroAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGyroAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorGyroAPIResponseModel).Reset()
}

// AlibabaInteractSensorGyroAPIResponseModel is 陀螺仪 成功返回结果
type AlibabaInteractSensorGyroAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_gyro_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0 表示正确
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorGyroAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorGyroAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorGyroAPIResponse)
	},
}

// GetAlibabaInteractSensorGyroAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorGyroAPIResponse
func GetAlibabaInteractSensorGyroAPIResponse() *AlibabaInteractSensorGyroAPIResponse {
	return poolAlibabaInteractSensorGyroAPIResponse.Get().(*AlibabaInteractSensorGyroAPIResponse)
}

// ReleaseAlibabaInteractSensorGyroAPIResponse 将 AlibabaInteractSensorGyroAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorGyroAPIResponse(v *AlibabaInteractSensorGyroAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorGyroAPIResponse.Put(v)
}
