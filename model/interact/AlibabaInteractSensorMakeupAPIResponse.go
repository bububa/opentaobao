package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorMakeupAPIResponse 美妆虚拟试装 API返回值
// alibaba.interact.sensor.makeup
//
// 手机淘宝美妆类目虚拟试妆权限，客户端能力（JS－API）
type AlibabaInteractSensorMakeupAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorMakeupAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorMakeupAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorMakeupAPIResponseModel).Reset()
}

// AlibabaInteractSensorMakeupAPIResponseModel is 美妆虚拟试装 成功返回结果
type AlibabaInteractSensorMakeupAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_makeup_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 非restAPI，为jsapi  result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorMakeupAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorMakeupAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorMakeupAPIResponse)
	},
}

// GetAlibabaInteractSensorMakeupAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorMakeupAPIResponse
func GetAlibabaInteractSensorMakeupAPIResponse() *AlibabaInteractSensorMakeupAPIResponse {
	return poolAlibabaInteractSensorMakeupAPIResponse.Get().(*AlibabaInteractSensorMakeupAPIResponse)
}

// ReleaseAlibabaInteractSensorMakeupAPIResponse 将 AlibabaInteractSensorMakeupAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorMakeupAPIResponse(v *AlibabaInteractSensorMakeupAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorMakeupAPIResponse.Put(v)
}
