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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_beneift_draw_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    ResultSuccess   bool `json:"result_success,omitempty" xml:"