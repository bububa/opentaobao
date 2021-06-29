package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销项目列表查询（已过时，不推荐使用） API返回值 
alibaba.damai.maitix.project.distribution.querylist

分销项目列表查询接口（已过时，不推荐使用）
*/
type AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixProjectDistributionQuerylistResponse
}

// 分销项目列表查询（已过时，不推荐使用） 成功返回结果
type AlibabaDamaiMaitixProjectDistributionQuerylistResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_project_distribution_querylist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 结果
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
