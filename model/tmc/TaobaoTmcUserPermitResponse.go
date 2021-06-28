package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
为已授权的用户开通消息服务 APIResponse
taobao.tmc.user.permit

为已授权的用户开通消息服务，授权消息使用。<br/><span style="color:red">注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic</span>
*/
type TaobaoTmcUserPermitAPIResponse struct {
    model.CommonResponse
    TaobaoTmcUserPermitResponse
}

type TaobaoTmcUserPermitResponse struct {
    XMLName xml.Name `xml:"tmc_user_permit_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
