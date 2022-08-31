package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse 根据商品id查询核销卷信息 API返回值
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
type AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel
}

// AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel is 根据商品id查询核销卷信息 成功返回结果
type AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_ticket_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectTicketQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
