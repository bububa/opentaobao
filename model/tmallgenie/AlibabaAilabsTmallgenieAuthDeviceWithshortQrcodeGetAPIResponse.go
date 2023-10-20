package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse 根据安全简码查询二维码详细信息 API返回值
// alibaba.ailabs.tmallgenie.auth.device.withshort.qrcode.get
//
// 根据安全简码查询二维码详细信息
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel is 根据安全简码查询二维码详细信息 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_withshort_qrcode_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 二维码数据
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 结果描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果码
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.Message = ""
	m.RetCode = 0
}

var poolAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceWithshortQrcodeGetAPIResponse.Put(v)
}
