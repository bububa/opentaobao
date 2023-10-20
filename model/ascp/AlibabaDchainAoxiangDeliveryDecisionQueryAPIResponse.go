package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangdeliverydecisionqueryAPIResponse 查询黑白名单快递 API返回值
// alibaba.dchain.aoxiang.delivery.decision.query
//
// 查询黑白名单快递
type AlibabadchainaoxiangdeliverydecisionqueryAPIResponse struct {
	model.CommonResponse
	AlibabadchainaoxiangdeliverydecisionqueryAPIResponseModel
}

// AlibabadchainaoxiangdeliverydecisionqueryAPIResponseModel is 查询黑白名单快递 成功返回结果
type AlibabadchainaoxiangdeliverydecisionqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_delivery_decision_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	DeliveryDecisionQueryResponse *DeliveryDecisionQueryResponse `json:"delivery_decision_query_response,omitempty" xml:"delivery_decision_query_response,omitempty"`
}
