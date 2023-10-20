package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAxWarehouseInboundCallbackAPIResponse 翱象入库回传 API返回值
// alibaba.ax.warehouse.inbound.callback
//
// 翱象入库回传
type AlibabaAxWarehouseInboundCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaAxWarehouseInboundCallbackAPIResponseModel
}

// AlibabaAxWarehouseInboundCallbackAPIResponseModel is 翱象入库回传 成功返回结果
type AlibabaAxWarehouseInboundCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ax_warehouse_inbound_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误编码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 错误信息
	ReturnMessage string `json:"return_message,omitempty" xml:"return_message,omitempty"`
	// 调用成功
	ReturnSuccess bool `json:"return_success,omitempty" xml:"return_success,omitempty"`
}
