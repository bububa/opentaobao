package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目分页查询项目列表服务 APIResponse
alibaba.damai.maitix.project.distribution.querybypage

分销项目分页查询项目列表服务
*/
type AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixProjectDistributionQuerybypageResponse
}

type AlibabaDamaiMaitixProjectDistributionQuerybypageResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_querybypage_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
