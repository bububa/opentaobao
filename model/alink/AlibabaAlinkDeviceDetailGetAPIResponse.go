package alink

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceDetailGetAPIResponse 获取设备详情 API返回值
// alibaba.alink.device.detail.get
//
// 阿里智能获取设备详情
type AlibabaAlinkDeviceDetailGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlinkDeviceDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlinkDeviceDetailGetAPIResponseModel).Reset()
}

// AlibabaAlinkDeviceDetailGetAPIResponseModel is 获取设备详情 成功返回结果
type AlibabaAlinkDeviceDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alink_device_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlinkDeviceDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlinkDeviceDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlinkDeviceDetailGetAPIResponse)
	},
}

// GetAlibabaAlinkDeviceDetailGetAPIResponse 从 sync.Pool 获取 AlibabaAlinkDeviceDetailGetAPIResponse
func GetAlibabaAlinkDeviceDetailGetAPIResponse() *AlibabaAlinkDeviceDetailGetAPIResponse {
	return poolAlibabaAlinkDeviceDetailGetAPIResponse.Get().(*AlibabaAlinkDeviceDetailGetAPIResponse)
}

// ReleaseAlibabaAlinkDeviceDetailGetAPIResponse 将 AlibabaAlinkDeviceDetailGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlinkDeviceDetailGetAPIResponse(v *AlibabaAlinkDeviceDetailGetAPIResponse) {
	v.Reset()
	poolAlibabaAlinkDeviceDetailGetAPIResponse.Put(v)
}
