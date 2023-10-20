package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse 根据mac查询设备的安全二维码 API返回值
// alibaba.ailabs.tmallgenie.auth.device.withmac.qrcode.get
//
// 根据mac查询二维码详细信息
type AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponseModel is 根据mac查询设备的安全二维码 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_withmac_qrcode_get_response"`
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
func (m *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
	m.Message = ""
	m.RetCode = 0
}

var poolAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceWithmacQrcodeGetAPIResponse.Put(v)
}
