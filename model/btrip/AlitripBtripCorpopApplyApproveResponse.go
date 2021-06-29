package btrip

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【商旅】更新审批单状态 APIResponse
alitrip.btrip.corpop.apply.approve

【商旅】更新审批单状态
*/
type AlitripBtripCorpopApplyApproveAPIResponse struct {
    model.CommonResponse
    AlitripBtripCorpopApplyApproveResponse
}

type AlitripBtripCorpopApplyApproveResponse struct {
    XMLName xml.Name `xml:"alitrip_btrip_corpop_apply_approve_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 传参
    
    Module   string `json:"module,omitempty" xml:"module,omitempty"`

    
    // 0
    
    ResultCode   int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 成功
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 成功标识
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
