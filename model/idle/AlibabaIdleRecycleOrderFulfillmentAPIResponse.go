package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRecycleOrderFulfillmentAPIResponse 闲鱼回收订单履约V1 API返回值
// alibaba.idle.recycle.order.fulfillment
//
// 外部回收商针对自有回收订单的履行
type AlibabaIdleRecycleOrderFulfillmentAPIResponse struct {
	model.CommonResponse
	AlibabaIdleRecycleOrderFulfillmentAPIResponseModel
}

// AlibabaIdleRecycleOrderFulfillmentAPIResponseModel is 闲鱼回收订单履约V1 成功返回结果
type AlibabaIdleRecycleOrderFulfillmentAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_fulfillment_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaIdleRecycleOrderFulfillmentResult `json:"result,omitempty" xml:"result,omitempty"`
}
