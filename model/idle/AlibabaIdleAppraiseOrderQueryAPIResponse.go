package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleAppraiseOrderQueryAPIResponse 闲鱼验货宝订单详情查询 API返回值
// alibaba.idle.appraise.order.query
//
// 鉴定商调用该接口获取订单状态
type AlibabaIdleAppraiseOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleAppraiseOrderQueryAPIResponseModel
}

// AlibabaIdleAppraiseOrderQueryAPIResponseModel is 闲鱼验货宝订单详情查询 成功返回结果
type AlibabaIdleAppraiseOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_appraise_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaIdleAppraiseOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
