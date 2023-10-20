package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAxWarehouseOutboundCallbackAPIResponse 翱象出仓回传 API返回值
// alibaba.ax.warehouse.outbound.callback
//
// 翱象出仓回传
type AlibabaAxWarehouseOutboundCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaAxWarehouseOutboundCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAxWarehouseOutboundCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAxWarehouseOutboundCallbackAPIResponseModel).Reset()
}

// AlibabaAxWarehouseOutboundCallbackAPIResponseModel is 翱象出仓回传 成功返回结果
type AlibabaAxWarehouseOutboundCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ax_warehouse_outbound_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误信息
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 调用成功
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAxWarehouseOutboundCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ReturnCode = ""
	m.ReturnMessage = ""
	m.ReturnSuccess = false
}

var poolAlibabaAxWarehouseOutboundCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAxWarehouseOutboundCallbackAPIResponse)
	},
}

// GetAlibabaAxWarehouseOutboundCallbackAPIResponse 从 sync.Pool 获取 AlibabaAxWarehouseOutboundCallbackAPIResponse
func GetAlibabaAxWarehouseOutboundCallbackAPIResponse() *AlibabaAxWarehouseOutboundCallbackAPIResponse {
	return poolAlibabaAxWarehouseOutboundCallbackAPIResponse.Get().(*AlibabaAxWarehouseOutboundCallbackAPIResponse)
}

// ReleaseAlibabaAxWarehouseOutboundCallbackAPIResponse 将 AlibabaAxWarehouseOutboundCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAxWarehouseOutboundCallbackAPIResponse(v *AlibabaAxWarehouseOutboundCallbackAPIResponse) {
	v.Reset()
	poolAlibabaAxWarehouseOutboundCallbackAPIResponse.Put(v)
}
