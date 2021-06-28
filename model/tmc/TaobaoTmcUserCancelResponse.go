package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取消用户的消息服务 APIResponse
taobao.tmc.user.cancel

取消用户的消息服务
*/
type TaobaoTmcUserCancelAPIResponse struct {
    model.CommonResponse
    TaobaoTmcUserCancelResponse
}

type TaobaoTmcUserCancelResponse struct {
    XMLName xml.Name `xml:"tmc_user_cancel_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功,如果为false并且没有错误码，表示删除的用户不存在。
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
