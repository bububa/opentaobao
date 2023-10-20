package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthnrceporderqueryAPIResponse 订单详情查询接口 API返回值
// alibaba.health.nr.cep.order.query
//
// 订单详情查询接口
type AlibabahealthnrceporderqueryAPIResponse struct {
	model.CommonResponse
	AlibabahealthnrceporderqueryAPIResponseModel
}

// AlibabahealthnrceporderqueryAPIResponseModel is 订单详情查询接口 成功返回结果
type AlibabahealthnrceporderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_health_nr_cep_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	ResponseResult *ResponseResult `json:"response_result,omitempty" xml:"response_result,omitempty"`
}
