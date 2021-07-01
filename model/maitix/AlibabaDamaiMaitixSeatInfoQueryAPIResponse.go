package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询座位信息 API返回值 
alibaba.damai.maitix.seat.info.query

分销查询座位文案信息
*/
type AlibabaDamaiMaitixSeatInfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel
}

// 分销商查询座位信息 成功返回结果
type AlibabaDamaiMaitixSeatInfoQueryAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_seat_info_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回信息
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
