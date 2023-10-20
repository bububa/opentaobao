package aligenie

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieDeviceUnbindAPIResponse 设备解绑操作接口 API返回值
// alibaba.ailabs.aligenie.device.unbind
//
// 让开发者能根据设备ID进行解绑操作的接口
type AlibabaAilabsAligenieDeviceUnbindAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsAligenieDeviceUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieDeviceUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabsAligenieDeviceUnbindAPIResponseModel).Reset()
}

// AlibabaAilabsAligenieDeviceUnbindAPIResponseModel is 设备解绑操作接口 成功返回结果
type AlibabaAilabsAligenieDeviceUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_aligenie_device_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 解绑是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabsAligenieDeviceUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaAilabsAligenieDeviceUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabsAligenieDeviceUnbindAPIResponse)
	},
}

// GetAlibabaAilabsAligenieDeviceUnbindAPIResponse 从 sync.Pool 获取 AlibabaAilabsAligenieDeviceUnbindAPIResponse
func GetAlibabaAilabsAligenieDeviceUnbindAPIResponse() *AlibabaAilabsAligenieDeviceUnbindAPIResponse {
	return poolAlibabaAilabsAligenieDeviceUnbindAPIResponse.Get().(*AlibabaAilabsAligenieDeviceUnbindAPIResponse)
}

// ReleaseAlibabaAilabsAligenieDeviceUnbindAPIResponse 将 AlibabaAilabsAligenieDeviceUnbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabsAligenieDeviceUnbindAPIResponse(v *AlibabaAilabsAligenieDeviceUnbindAPIResponse) {
	v.Reset()
	poolAlibabaAilabsAligenieDeviceUnbindAPIResponse.Put(v)
}
