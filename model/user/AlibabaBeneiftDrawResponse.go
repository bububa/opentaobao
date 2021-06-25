package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖接口 APIResponse
alibaba.beneift.draw

抽奖接口
*/
type AlibabaBeneiftDrawAPIResponse struct {
    model.CommonResponse
    Response *AlibabaBeneiftDrawResponse `json:"alibaba_beneift_draw_response,omitempty"`
}

type AlibabaBeneiftDrawResponse struct {

    // 是否成功
    ResultSuccess   bool `json:"result_success,omitempty"`

    // message
    ResultMsg   string `json:"result_msg,omitempty"`

    // code
    ResultCode   string `json:"result_code,omitempty"`

    // 权益id
    RightId   string `json:"right_id,omitempty"`

    // 奖品id
    PrizeId   string `json:"prize_id,omitempty"`

}
