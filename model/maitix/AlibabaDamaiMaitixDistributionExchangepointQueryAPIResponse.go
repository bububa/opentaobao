package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销查询取票点接口 API返回值 
alibaba.damai.maitix.distribution.exchangepoint.query

分销查询取票点接口
*/
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel
}

// 分销查询取票点接口 成功返回结果
type AlibabaDamaiMaitixDistributionExchangepointQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_distribution_exchangepoint_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
