package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销查询取票点接口 APIResponse
alibaba.damai.maitix.distribution.exchangepoint.query

分销查询取票点接口
*/
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixDistributionExchangepointQueryResponse
}

type AlibabaDamaiMaitixDistributionExchangepointQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_exchangepoint_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
