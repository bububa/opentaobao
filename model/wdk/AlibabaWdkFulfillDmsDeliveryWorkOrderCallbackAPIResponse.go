package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse 末端配配送作业回传 API返回值
// alibaba.wdk.fulfill.dms.delivery.work.order.callback
//
// 末端配配送作业回传。
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponseModel).Reset()
}

// AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponseModel is 末端配配送作业回传 成功返回结果
type AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_dms_delivery_work_order_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应提示信息
	RespMessage string `json:"resp_message,omitempty" xml:"resp_message,omitempty"`
	// 返回码： SUCCESS:作业批次接收成功 SYSTEM_ERROR :系统异常(指令可重发) PARAM_ERROR :参数错误(指令不可重发，监控报警) BUSINESS_ERROR:业务异常(指令不可重发，监控报警)
	RespCode string `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 是否成功：true 调用成功； false 调用失败
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RespMessage = ""
	m.RespCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse)
	},
}

// GetAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse
func GetAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse() *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse {
	return poolAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse.Get().(*AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse)
}

// ReleaseAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse 将 AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse(v *AlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillDmsDeliveryWorkOrderCallbackAPIResponse.Put(v)
}
