package iot

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDevicePayUrlGetAPIResponse 贩卖机支付二维链接获取 API返回值
// alibaba.retail.device.payUrl.get
//
// 贩卖机支付二维链接获取
type AlibabaRetailDevicePayUrlGetAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDevicePayUrlGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailDevicePayUrlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailDevicePayUrlGetAPIResponseModel).Reset()
}

// AlibabaRetailDevicePayUrlGetAPIResponseModel is 贩卖机支付二维链接获取 成功返回结果
type AlibabaRetailDevicePayUrlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_payUrl_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaRetailDevicePayUrlGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailDevicePayUrlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailDevicePayUrlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDevicePayUrlGetAPIResponse)
	},
}

// GetAlibabaRetailDevicePayUrlGetAPIResponse 从 sync.Pool 获取 AlibabaRetailDevicePayUrlGetAPIResponse
func GetAlibabaRetailDevicePayUrlGetAPIResponse() *AlibabaRetailDevicePayUrlGetAPIResponse {
	return poolAlibabaRetailDevicePayUrlGetAPIResponse.Get().(*AlibabaRetailDevicePayUrlGetAPIResponse)
}

// ReleaseAlibabaRetailDevicePayUrlGetAPIResponse 将 AlibabaRetailDevicePayUrlGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailDevicePayUrlGetAPIResponse(v *AlibabaRetailDevicePayUrlGetAPIResponse) {
	v.Reset()
	poolAlibabaRetailDevicePayUrlGetAPIResponse.Put(v)
}
