package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse 设备状态查询 API返回值
// alibaba.ailabs.tmallgenie.auth.device.status.get
//
// 提供给天猫精灵定制机厂商 查询设备在线状态值
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel is 设备状态查询 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口对象封装
	Result *AlibabaAilabsTmallgenieAuthDeviceStatusGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIResponse.Put(v)
}
