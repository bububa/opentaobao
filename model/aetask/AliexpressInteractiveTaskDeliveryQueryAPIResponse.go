package aetask

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressInteractiveTaskDeliveryQueryAPIResponse AE互动任务投放 API返回值
// aliexpress.interactive.task.delivery.query
//
// 将内部配置好的任务，如浏览商品，店铺投放给外部ISV
type AliexpressInteractiveTaskDeliveryQueryAPIResponse struct {
	model.CommonResponse
	AliexpressInteractiveTaskDeliveryQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressInteractiveTaskDeliveryQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressInteractiveTaskDeliveryQueryAPIResponseModel).Reset()
}

// AliexpressInteractiveTaskDeliveryQueryAPIResponseModel is AE互动任务投放 成功返回结果
type AliexpressInteractiveTaskDeliveryQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_interactive_task_delivery_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务返回接口
	Results []AliexpressInteractiveTaskDeliveryQueryResult `json:"results,omitempty" xml:"results>aliexpress_interactive_task_delivery_query_result,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressInteractiveTaskDeliveryQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.MsgInfo = ""
	m.MsgCode = ""
}

var poolAliexpressInteractiveTaskDeliveryQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressInteractiveTaskDeliveryQueryAPIResponse)
	},
}

// GetAliexpressInteractiveTaskDeliveryQueryAPIResponse 从 sync.Pool 获取 AliexpressInteractiveTaskDeliveryQueryAPIResponse
func GetAliexpressInteractiveTaskDeliveryQueryAPIResponse() *AliexpressInteractiveTaskDeliveryQueryAPIResponse {
	return poolAliexpressInteractiveTaskDeliveryQueryAPIResponse.Get().(*AliexpressInteractiveTaskDeliveryQueryAPIResponse)
}

// ReleaseAliexpressInteractiveTaskDeliveryQueryAPIResponse 将 AliexpressInteractiveTaskDeliveryQueryAPIResponse 保存到 sync.Pool
func ReleaseAliexpressInteractiveTaskDeliveryQueryAPIResponse(v *AliexpressInteractiveTaskDeliveryQueryAPIResponse) {
	v.Reset()
	poolAliexpressInteractiveTaskDeliveryQueryAPIResponse.Put(v)
}
