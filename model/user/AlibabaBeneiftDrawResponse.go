package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖接口 APIResponse
alibaba.beneift.draw

抽奖接口
*/
type AlibabaBeneiftDrawAPIResponse struct {
    model.CommonResponse
    AlibabaBeneiftDrawResponse
}

type AlibabaBeneiftDrawResponse struct {
    XMLName xml.Name `xml:"alibaba_beneift_draw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"result_success,omitempty"`

    
    // message
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 权益id
    
    RightId   string `json:"right_id,omitempty" xml:"right_id,omitempty"`

    
    // 奖品id
    
    PrizeId   string `json:"prize_id,omitempty" xml:"prize_id,omitempty"`

    
}
