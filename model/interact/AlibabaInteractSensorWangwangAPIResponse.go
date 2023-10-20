package interact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorWangwangAPIResponse 手淘拉起旺旺接口 API返回值
// alibaba.interact.sensor.wangwang
//
// 手淘开放专用接口，没有数据返回，仅用于手淘容器中jssdk接口鉴权。手淘开放旺旺拉起功能给ISV。
type AlibabaInteractSensorWangwangAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorWangwangAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorWangwangAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorWangwangAPIResponseModel).Reset()
}

// AlibabaInteractSensorWangwangAPIResponseModel is 手淘拉起旺旺接口 成功返回结果
type AlibabaInteractSensorWangwangAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_wangwang_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=0
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorWangwangAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorWangwangAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorWangwangAPIResponse)
	},
}

// GetAlibabaInteractSensorWangwangAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorWangwangAPIResponse
func GetAlibabaInteractSensorWangwangAPIResponse() *AlibabaInteractSensorWangwangAPIResponse {
	return poolAlibabaInteractSensorWangwangAPIResponse.Get().(*AlibabaInteractSensorWangwangAPIResponse)
}

// ReleaseAlibabaInteractSensorWangwangAPIResponse 将 AlibabaInteractSensorWangwangAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorWangwangAPIResponse(v *AlibabaInteractSensorWangwangAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorWangwangAPIResponse.Put(v)
}
