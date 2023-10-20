package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse 根据三方ID查询设备注册激活信息 API返回值
// alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get
//
// 根据三方ID查询设备注册激活信息
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel is 根据三方ID查询设备注册激活信息 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_withdeviceid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse.Put(v)
}
