package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvOrderQueryAPIResponse
闲鱼已验货订单查询 API返回值
alibaba.idle.isv.order.query

服务商ISV，根据订单号，查询闲鱼订单信息 */
type AlibabaIdleIsvOrderQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvOrderQueryAPIResponseModel
}

// AlibabaIdleIsvOrderQueryAPIResponseModel is 闲鱼已验货订单查询 成功返回结果
type AlibabaIdleIsvOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleIsvOrderQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
