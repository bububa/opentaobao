package ticket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripticketproductqueryAPIResponse 【门票API2.0】门票商品查询接口 API返回值
// alitrip.ticket.product.query
//
// 门票商品查询接口：返回商家上传的门票商品信息
type AlitripticketproductqueryAPIResponse struct {
	model.CommonResponse
	AlitripticketproductqueryAPIResponseModel
}

// AlitripticketproductqueryAPIResponseModel is 【门票API2.0】门票商品查询接口 成功返回结果
type AlitripticketproductqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ticket_product_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 门票商品详情
	FirstResult *TopTicketItemFullinfoResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
