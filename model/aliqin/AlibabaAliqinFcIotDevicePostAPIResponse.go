package aliqin

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotDevicePostAPIResponse 商家提交设备信息 API返回值
// alibaba.aliqin.fc.iot.device.post
//
// 物联网商家设备信息录入
type AlibabaAliqinFcIotDevicePostAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotDevicePostAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotDevicePostAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinFcIotDevicePostAPIResponseModel).Reset()
}

// AlibabaAliqinFcIotDevicePostAPIResponseModel is 商家提交设备信息 成功返回结果
type AlibabaAliqinFcIotDevicePostAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_device_post_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotDevicePostResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinFcIotDevicePostAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinFcIotDevicePostAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinFcIotDevicePostAPIResponse)
	},
}

// GetAlibabaAliqinFcIotDevicePostAPIResponse 从 sync.Pool 获取 AlibabaAliqinFcIotDevicePostAPIResponse
func GetAlibabaAliqinFcIotDevicePostAPIResponse() *AlibabaAliqinFcIotDevicePostAPIResponse {
	return poolAlibabaAliqinFcIotDevicePostAPIResponse.Get().(*AlibabaAliqinFcIotDevicePostAPIResponse)
}

// ReleaseAlibabaAliqinFcIotDevicePostAPIResponse 将 AlibabaAliqinFcIotDevicePostAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinFcIotDevicePostAPIResponse(v *AlibabaAliqinFcIotDevicePostAPIResponse) {
	v.Reset()
	poolAlibabaAliqinFcIotDevicePostAPIResponse.Put(v)
}
