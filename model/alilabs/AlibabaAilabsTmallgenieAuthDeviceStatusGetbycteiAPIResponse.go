package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse 根据ctei查询状态 API返回值
// alibaba.ailabs.tmallgenie.auth.device.status.getbyctei
//
// 提供给电信查询设备在线状态值
type AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponseModel is 根据ctei查询状态 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_status_getbyctei_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口对象封装
	Result *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceStatusGetbycteiAPIResponse.Put(v)
}
