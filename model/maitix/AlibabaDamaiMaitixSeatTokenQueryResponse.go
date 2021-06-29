package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商选座获取qtoken API返回值 
alibaba.damai.maitix.seat.token.query

选座分销，分销商查询token
*/
type AlibabaDamaiMaitixSeatTokenQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixSeatTokenQueryResponse
}

// 分销商选座获取qtoken 成功返回结果
type AlibabaDamaiMaitixSeatTokenQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_seat_token_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 系统自动生成
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}