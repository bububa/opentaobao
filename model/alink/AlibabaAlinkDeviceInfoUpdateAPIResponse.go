package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceInfoUpdateAPIResponse 更新设备昵称等信息 API返回值
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
type AlibabaAlinkDeviceInfoUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceInfoUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceInfoUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkDeviceInfoUpdateAPIResponseModel).Reset()
}

// AlibabaAlinkDeviceInfoUpdateAPIResponseModel is 更新设备昵称等信息 成功返回结果
type AlibabaAlinkDeviceInfoUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_info_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceInfoUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkDeviceInfoUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkDeviceInfoUpdateAPIResponse)
	},
}

// GetAlibabaAlinkDeviceInfoUpdateAPIResponse 从 sync.Pool 获取 AlibabaAlinkDeviceInfoUpdateAPIResponse
func GetAlibabaAlinkDeviceInfoUpdateAPIResponse() *AlibabaAlinkDeviceInfoUpdateAPIResponse {
	return poolAlibabaAlinkDeviceInfoUpdateAPIResponse.Get().(*AlibabaAlinkDeviceInfoUpdateAPIResponse)
}

// ReleaseAlibabaAlinkDeviceInfoUpdateAPIResponse 将 AlibabaAlinkDeviceInfoUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkDeviceInfoUpdateAPIResponse(v *AlibabaAlinkDeviceInfoUpdateAPIResponse) {
	v.Reset()
	poolAlibabaAlinkDeviceInfoUpdateAPIResponse.Put(v)
}
