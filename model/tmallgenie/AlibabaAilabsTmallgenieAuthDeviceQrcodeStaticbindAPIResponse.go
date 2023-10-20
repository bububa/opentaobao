package tmallgenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse 静态二维码绑定 API返回值
// alibaba.ailabs.tmallgenie.auth.device.qrcode.staticbind
//
// 静态二维码绑定
type AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel).Reset()
}

// AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel is 静态二维码绑定 成功返回结果
type AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_auth_device_qrcode_staticbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse)
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse
func GetAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse() *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse {
	return poolAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse.Get().(*AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse 将 AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse(v *AlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceQrcodeStaticbindAPIResponse.Put(v)
}
