package maitix

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商查询座位信息 APIResponse
alibaba.damai.maitix.seat.info.query

分销查询座位文案信息
*/
type AlibabaDamaiMaitixSeatInfoQueryAPIResponse struct {
    model.CommonResponse
    AlibabaDamaiMaitixSeatInfoQueryResponse
}

type AlibabaDamaiMaitixSeatInfoQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_damai_maitix_seat_info_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回信息
    
    Result   *OpenResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
