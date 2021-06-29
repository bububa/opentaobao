package flightuppc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
航司权益数据回流 APIResponse
alitrip.uppc.member.gain

航司权益数据回流
*/
type AlitripUppcMemberGainAPIResponse struct {
    model.CommonResponse
    AlitripUppcMemberGainResponse
}

type AlitripUppcMemberGainResponse struct {
    XMLName xml.Name `xml:"alitrip_uppc_member_gain_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *ResultDO `json:"result,omitempty" xml:"result,omitempty"`

    
}
