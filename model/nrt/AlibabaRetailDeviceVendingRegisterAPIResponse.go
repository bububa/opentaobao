package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDeviceVendingRegisterAPIResponse 贩卖机设备注册 API返回值
// alibaba.retail.device.vending.register
//
// 贩卖机注册
type AlibabaRetailDeviceVendingRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaRetailDeviceVendingRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceVendingRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailDeviceVendingRegisterAPIResponseModel).Reset()
}

// AlibabaRetailDeviceVendingRegisterAPIResponseModel is 贩卖机设备注册 成功返回结果
type AlibabaRetailDeviceVendingRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_device_vending_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaRetailDeviceVendingRegisterResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailDeviceVendingRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailDeviceVendingRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailDeviceVendingRegisterAPIResponse)
	},
}

// GetAlibabaRetailDeviceVendingRegisterAPIResponse 从 sync.Pool 获取 AlibabaRetailDeviceVendingRegisterAPIResponse
func GetAlibabaRetailDeviceVendingRegisterAPIResponse() *AlibabaRetailDeviceVendingRegisterAPIResponse {
	return poolAlibabaRetailDeviceVendingRegisterAPIResponse.Get().(*AlibabaRetailDeviceVendingRegisterAPIResponse)
}

// ReleaseAlibabaRetailDeviceVendingRegisterAPIResponse 将 AlibabaRetailDeviceVendingRegisterAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailDeviceVendingRegisterAPIResponse(v *AlibabaRetailDeviceVendingRegisterAPIResponse) {
	v.Reset()
	poolAlibabaRetailDeviceVendingRegisterAPIResponse.Put(v)
}
