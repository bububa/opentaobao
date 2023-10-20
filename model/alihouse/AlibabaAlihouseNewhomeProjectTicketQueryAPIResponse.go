package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectticketqueryAPIResponse 根据商品id查询核销卷信息 API返回值
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
type AlibabaalihousenewhomeprojectticketqueryAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomeprojectticketqueryAPIResponseModel
}

// AlibabaalihousenewhomeprojectticketqueryAPIResponseModel is 根据商品id查询核销卷信息 成功返回结果
type AlibabaalihousenewhomeprojectticketqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_ticket_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomeprojectticketqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
