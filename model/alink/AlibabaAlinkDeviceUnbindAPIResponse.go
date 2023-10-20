package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceUnbindAPIResponse 解绑设备 API返回值
// alibaba.alink.device.unbind
//
// 阿里智能解绑设备
type AlibabaAlinkDeviceUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkDeviceUnbindAPIResponseModel).Reset()
}

// AlibabaAlinkDeviceUnbindAPIResponseModel is 解绑设备 成功返回结果
type AlibabaAlinkDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkDeviceUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkDeviceUnbindAPIResponse)
	},
}

// GetAlibabaAlinkDeviceUnbindAPIResponse 从 sync.Pool 获取 AlibabaAlinkDeviceUnbindAPIResponse
func GetAlibabaAlinkDeviceUnbindAPIResponse() *AlibabaAlinkDeviceUnbindAPIResponse {
	return poolAlibabaAlinkDeviceUnbindAPIResponse.Get().(*AlibabaAlinkDeviceUnbindAPIResponse)
}

// ReleaseAlibabaAlinkDeviceUnbindAPIResponse 将 AlibabaAlinkDeviceUnbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkDeviceUnbindAPIResponse(v *AlibabaAlinkDeviceUnbindAPIResponse) {
	v.Reset()
	poolAlibabaAlinkDeviceUnbindAPIResponse.Put(v)
}
