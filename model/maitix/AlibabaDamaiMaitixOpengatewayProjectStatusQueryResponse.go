package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销状态查询接口queryProjectStatusByProjectId APIResponse
alibaba.damai.maitix.opengateway.project.status.query

queryProjectStatusByProjectId
*/
type AlibabaDamaiMaitixOpengatewayProjectStatusQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixOpengatewayProjectStatusQueryResponse
}

type AlibabaDamaiMaitixOpengatewayProjectStatusQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_opengateway_project_status_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
