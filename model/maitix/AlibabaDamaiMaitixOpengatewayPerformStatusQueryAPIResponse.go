package maitix

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabadamaimaitixopengatewayperformstatusqueryAPIResponse 分销状态查询接口queryPerformStatusByPerformId API返回值
// alibaba.damai.maitix.opengateway.perform.status.query
//
// queryPerformStatusByPerformId
type AlibabadamaimaitixopengatewayperformstatusqueryAPIResponse struct {
	model.CommonResponse
	AlibabadamaimaitixopengatewayperformstatusqueryAPIResponseModel
}

// AlibabadamaimaitixopengatewayperformstatusqueryAPIResponseModel is 分销状态查询接口queryPerformStatusByPerformId 成功返回结果
type AlibabadamaimaitixopengatewayperformstatusqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_perform_status_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
