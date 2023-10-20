package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceDeviceOrderUpdateAPIResponse 回传/更新设备订购单 API返回值
// alibaba.einvoice.device.order.update
//
// 更新设备订购单，同步税控设备信息
type AlibabaEinvoiceDeviceOrderUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceDeviceOrderUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceDeviceOrderUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceDeviceOrderUpdateAPIResponseModel).Reset()
}

// AlibabaEinvoiceDeviceOrderUpdateAPIResponseModel is 回传/更新设备订购单 成功返回结果
type AlibabaEinvoiceDeviceOrderUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_device_order_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceDeviceOrderUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceDeviceOrderUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceDeviceOrderUpdateAPIResponse)
	},
}

// GetAlibabaEinvoiceDeviceOrderUpdateAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceDeviceOrderUpdateAPIResponse
func GetAlibabaEinvoiceDeviceOrderUpdateAPIResponse() *AlibabaEinvoiceDeviceOrderUpdateAPIResponse {
	return poolAlibabaEinvoiceDeviceOrderUpdateAPIResponse.Get().(*AlibabaEinvoiceDeviceOrderUpdateAPIResponse)
}

// ReleaseAlibabaEinvoiceDeviceOrderUpdateAPIResponse 将 AlibabaEinvoiceDeviceOrderUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceDeviceOrderUpdateAPIResponse(v *AlibabaEinvoiceDeviceOrderUpdateAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceDeviceOrderUpdateAPIResponse.Put(v)
}
