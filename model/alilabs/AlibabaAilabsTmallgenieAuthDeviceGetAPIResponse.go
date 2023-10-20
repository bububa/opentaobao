package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse 获取设备详情 API返回值
// alibaba.ailabs.tmallgenie.auth.device.get
//
// 通过此接口获取设备详情
type AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceGetAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceGetAPIResponseModel is 获取设备详情 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备信息
	DeviceInfo *AlibabaAilabsTmallgenieAuthDeviceGetResult `json:"device_info,omitempty" xml:"device_info,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeviceInfo = nil
}

var poolAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceGetAPIResponse.Put(v)
}
