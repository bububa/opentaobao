package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse 解绑设备 API返回值
// alibaba.ailabs.tmallgenie.auth.device.unbind
//
// 通过此接口解绑天猫精灵设备
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel is 解绑设备 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *AlibabaAilabsTmallgenieAuthDeviceUnbindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse.Put(v)
}
