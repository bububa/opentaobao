package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取mall的离线等级 APIResponse
alibaba.mos.ark.health.offline.level.get

获取mall的离线等级
*/
type AlibabaMosArkHealthOfflineLevelGetAPIResponse struct {
    model.CommonResponse
    AlibabaMosArkHealthOfflineLevelGetResponse
}

type AlibabaMosArkHealthOfflineLevelGetResponse struct {
    XMLName xml.Name `xml:"alibaba_mos_ark_health_offline_level_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // model
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // msgInfo
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
}
