package flightuppc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
航司权益数据回流 API返回值 
alitrip.uppc.member.gain

航司权益数据回流
*/
type AlitripUppcMemberGainAPIResponse struct {
    model.CommonResponse
    AlitripUppcMemberGainAPIResponseModel
}

// 航司权益数据回流 成功返回结果
type AlitripUppcMemberGainAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_uppc_member_gain_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
