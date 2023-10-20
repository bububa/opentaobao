package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceBindAPIResponse 绑定设备 API返回值
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
type AlibabaAlinkDeviceBindAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceBindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkDeviceBindAPIResponseModel).Reset()
}

// AlibabaAlinkDeviceBindAPIResponseModel is 绑定设备 成功返回结果
type AlibabaAlinkDeviceBindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkDeviceBindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkDeviceBindAPIResponse)
	},
}

// GetAlibabaAlinkDeviceBindAPIResponse 从 sync.Pool 获取 AlibabaAlinkDeviceBindAPIResponse
func GetAlibabaAlinkDeviceBindAPIResponse() *AlibabaAlinkDeviceBindAPIResponse {
	return poolAlibabaAlinkDeviceBindAPIResponse.Get().(*AlibabaAlinkDeviceBindAPIResponse)
}

// ReleaseAlibabaAlinkDeviceBindAPIResponse 将 AlibabaAlinkDeviceBindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkDeviceBindAPIResponse(v *AlibabaAlinkDeviceBindAPIResponse) {
	v.Reset()
	poolAlibabaAlinkDeviceBindAPIResponse.Put(v)
}
