package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发奖接口 APIResponse
alibaba.benefit.send

发奖接口
*/
type AlibabaBenefitSendAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_benefit_send_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 接口返回消息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"