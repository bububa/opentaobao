package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据互动实例ID查询奖池信息 APIResponse
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息
*/
type AlibabaInteractIsvadminGetpondbyinteractAPIResponse struct {
    model.CommonResponse
    AlibabaInteractIsvadminGetpondbyinteractResponse
}

type AlibabaInteractIsvadminGetpondbyinteractResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_isvadmin_getpondbyinteract_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 奖池信息
    
    Data   *PrizePondDTO `json:"data,omitempty" xml:"data,omitempty"`

    
    // 是否调用成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
    // 调用错误原因
    
    Error   string `json:"error,omitempty" xml:"error,omitempty"`

    
    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`

    
}
