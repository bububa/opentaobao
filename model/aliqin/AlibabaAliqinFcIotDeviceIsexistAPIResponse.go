package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotDeviceIsexistAPIResponse 判断设备是否存在 API返回值
// alibaba.aliqin.fc.iot.device.isexist
//
// 判断设备是否存在
type AlibabaAliqinFcIotDeviceIsexistAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotDeviceIsexistAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotDeviceIsexistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotDeviceIsexistAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotDeviceIsexistAPIResponseModel is 判断设备是否存在 成功返回结果
type AlibabaAliqinFcIotDeviceIsexistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_device_isexist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotDeviceIsexistResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotDeviceIsexistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotDeviceIsexistAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotDeviceIsexistAPIResponse)
	},
}

// GetAlibabaAliqinFcIotDeviceIsexistAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotDeviceIsexistAPIResponse
func GetAlibabaAliqinFcIotDeviceIsexistAPIResponse() *AlibabaAliqinFcIotDeviceIsexistAPIResponse {
	return poolAlibabaAliqinFcIotDeviceIsexistAPIResponse.Get().(*AlibabaAliqinFcIotDeviceIsexistAPIResponse)
}

// ReleaseAlibabaAliqinFcIotDeviceIsexistAPIResponse 将 AlibabaAliqinFcIotDeviceIsexistAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotDeviceIsexistAPIResponse(v *AlibabaAliqinFcIotDeviceIsexistAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotDeviceIsexistAPIResponse.Put(v)
}
