package aetask

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressInteractiveTaskDeliveryQueryAPIResponse
AE互动任务投放 API返回值
aliexpress.interactive.task.delivery.query

将内部配置好的任务，如浏览商品，店铺投放给外部ISV */
type AliexpressInteractiveTaskDeliveryQueryAPIResponse struct {
	model.CommonResponse
	AliexpressInteractiveTaskDeliveryQueryAPIResponseModel
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
