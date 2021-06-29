package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryPerformStatusByPerformId APIResponse
alibaba.damai.maitix.opengateway.perform.status.query

queryPerformStatusByPerformId
*/
type AlibabaDamaiMaitixOpengatewayPerformStatusQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixOpengatewayPerformStatusQueryResponse
}

type AlibabaDamaiMaitixOpengatewayPerformStatusQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_perform_status_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 查询结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
