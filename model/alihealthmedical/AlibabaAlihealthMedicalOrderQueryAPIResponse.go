package alihealthmedical

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmedicalorderqueryAPIResponse 三方机构查询订单详情接口 API返回值
// alibaba.alihealth.medical.order.query
//
// 查询订单详情，包括评价
type AlibabaalihealthmedicalorderqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthmedicalorderqueryAPIResponseModel
}

// AlibabaalihealthmedicalorderqueryAPIResponseModel is 三方机构查询订单详情接口 成功返回结果
type AlibabaalihealthmedicalorderqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_medical_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	ServiceResult *ServiceResult `json:"service_result,omitempty" xml:"service_result,omitempty"`
}
