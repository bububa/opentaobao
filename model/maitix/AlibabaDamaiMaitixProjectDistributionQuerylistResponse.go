package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目列表查询（已过时，不推荐使用） APIResponse
alibaba.damai.maitix.project.distribution.querylist

分销项目列表查询接口（已过时，不推荐使用）
*/
type AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixProjectDistributionQuerylistResponse
}

type AlibabaDamaiMaitixProjectDistributionQuerylistResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_querylist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
