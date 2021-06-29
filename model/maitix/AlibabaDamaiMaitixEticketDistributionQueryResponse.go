package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销电子票查询接口 APIResponse
alibaba.damai.maitix.eticket.distribution.query

分销电子票查询接口
*/
type AlibabaDamaiMaitixEticketDistributionQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixEticketDistributionQueryResponse
}

type AlibabaDamaiMaitixEticketDistributionQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_eticket_distribution_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
