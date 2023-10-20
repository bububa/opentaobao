package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeDeviceReturnAPIResponse 服务商回传客户端设备列表 API返回值
// alibaba.einvoice.income.device.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
type AlibabaEinvoiceIncomeDeviceReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceIncomeDeviceReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceIncomeDeviceReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceIncomeDeviceReturnAPIResponseModel).Reset()
}

// AlibabaEinvoiceIncomeDeviceReturnAPIResponseModel is 服务商回传客户端设备列表 成功返回结果
type AlibabaEinvoiceIncomeDeviceReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_income_device_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口是否调用成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceIncomeDeviceReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceIncomeDeviceReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceIncomeDeviceReturnAPIResponse)
	},
}

// GetAlibabaEinvoiceIncomeDeviceReturnAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceIncomeDeviceReturnAPIResponse
func GetAlibabaEinvoiceIncomeDeviceReturnAPIResponse() *AlibabaEinvoiceIncomeDeviceReturnAPIResponse {
	return poolAlibabaEinvoiceIncomeDeviceReturnAPIResponse.Get().(*AlibabaEinvoiceIncomeDeviceReturnAPIResponse)
}

// ReleaseAlibabaEinvoiceIncomeDeviceReturnAPIResponse 将 AlibabaEinvoiceIncomeDeviceReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceIncomeDeviceReturnAPIResponse(v *AlibabaEinvoiceIncomeDeviceReturnAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceIncomeDeviceReturnAPIResponse.Put(v)
}
