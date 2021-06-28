package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取天猫互动奖池列表 APIResponse
alibaba.interact.isvadmin.allponds

获取天猫互动奖池列表
*/
type AlibabaInteractIsvadminAllpondsAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvadminAllpondsResponse
}

type AlibabaInteractIsvadminAllpondsResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_isvadmin_allponds_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 错误码
    
    InteractErrorCode   string `json:"interact_error_code,omitempty" xml:"interact_error_code,omitempty"`

    
    // 错误描述
    
    InteractErrorMsg   string `json:"interact_error_msg,omitempty" xml:"interact_error_msg,omitempty"`

    
    // 奖池列表
    
    Allponds   []PrizePondDTO `json:"allponds,omitempty" xml:"allponds>prize_pond_dto,omitempty"`
    
    
}
