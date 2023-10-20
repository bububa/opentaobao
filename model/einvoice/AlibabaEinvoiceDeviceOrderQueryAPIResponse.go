package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceDeviceOrderQueryAPIResponse 查询税控设备加盘订购单详情 API返回值
// alibaba.einvoice.device.order.query
//
// 查询税控设备订购单详情
type AlibabaEinvoiceDeviceOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceDeviceOrderQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceDeviceOrderQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceDeviceOrderQueryAPIResponseModel).Reset()
}

// AlibabaEinvoiceDeviceOrderQueryAPIResponseModel is 查询税控设备加盘订购单详情 成功返回结果
type AlibabaEinvoiceDeviceOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_device_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceDeviceOrderQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaEinvoiceDeviceOrderQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceDeviceOrderQueryAPIResponse)
	},
}

// GetAlibabaEinvoiceDeviceOrderQueryAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceDeviceOrderQueryAPIResponse
func GetAlibabaEinvoiceDeviceOrderQueryAPIResponse() *AlibabaEinvoiceDeviceOrderQueryAPIResponse {
	return poolAlibabaEinvoiceDeviceOrderQueryAPIResponse.Get().(*AlibabaEinvoiceDeviceOrderQueryAPIResponse)
}

// ReleaseAlibabaEinvoiceDeviceOrderQueryAPIResponse 将 AlibabaEinvoiceDeviceOrderQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceDeviceOrderQueryAPIResponse(v *AlibabaEinvoiceDeviceOrderQueryAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceDeviceOrderQueryAPIResponse.Put(v)
}
