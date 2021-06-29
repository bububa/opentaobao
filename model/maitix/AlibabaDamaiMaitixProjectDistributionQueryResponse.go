package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销单个项目信息查询 APIResponse
alibaba.damai.maitix.project.distribution.query

发布分销项目查询单个项目信息接口
*/
type AlibabaDamaiMaitixProjectDistributionQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixProjectDistributionQueryResponse
}

type AlibabaDamaiMaitixProjectDistributionQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
