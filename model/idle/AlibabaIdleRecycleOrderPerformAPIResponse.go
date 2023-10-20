package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerecycleorderperformAPIResponse 回收订单履约V2 API返回值
// alibaba.idle.recycle.order.perform
//
// 闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
type AlibabaidlerecycleorderperformAPIResponse struct {
	model.CommonResponse
	AlibabaidlerecycleorderperformAPIResponseModel
}

// AlibabaidlerecycleorderperformAPIResponseModel is 回收订单履约V2 成功返回结果
type AlibabaidlerecycleorderperformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_recycle_order_perform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
