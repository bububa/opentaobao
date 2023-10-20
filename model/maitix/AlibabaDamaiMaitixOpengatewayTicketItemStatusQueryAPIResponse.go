package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponse 分销状态查询接口queryTicketItemStatusByTicketItemId API返回值
// alibaba.damai.maitix.opengateway.ticketItem.status.query
//
// queryTicketItemStatusByTicketItemId
type AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponse struct {
	model.CommonResponse
	AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponseModel
}

// AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponseModel is 分销状态查询接口queryTicketItemStatusByTicketItemId 成功返回结果
type AlibabadamaimaitixopengatewayticketItemstatusqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_ticketItem_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
