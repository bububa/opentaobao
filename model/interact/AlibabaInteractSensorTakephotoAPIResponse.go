package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTakephotoAPIResponse takePhoto API返回值
// alibaba.interact.sensor.takephoto
//
// 客户端takePhoto
type AlibabaInteractSensorTakephotoAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorTakephotoAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorTakephotoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorTakephotoAPIResponseModel).Reset()
}

// AlibabaInteractSensorTakephotoAPIResponseModel is takePhoto 成功返回结果
type AlibabaInteractSensorTakephotoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_takephoto_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// return=0表示成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorTakephotoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorTakephotoAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorTakephotoAPIResponse)
	},
}

// GetAlibabaInteractSensorTakephotoAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorTakephotoAPIResponse
func GetAlibabaInteractSensorTakephotoAPIResponse() *AlibabaInteractSensorTakephotoAPIResponse {
	return poolAlibabaInteractSensorTakephotoAPIResponse.Get().(*AlibabaInteractSensorTakephotoAPIResponse)
}

// ReleaseAlibabaInteractSensorTakephotoAPIResponse 将 AlibabaInteractSensorTakephotoAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorTakephotoAPIResponse(v *AlibabaInteractSensorTakephotoAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorTakephotoAPIResponse.Put(v)
}
