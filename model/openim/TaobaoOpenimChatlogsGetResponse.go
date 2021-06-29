package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openim聊天记录查询接口 APIResponse
taobao.openim.chatlogs.get

查询openim账号聊天记录
*/
type TaobaoOpenimChatlogsGetAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimChatlogsGetResponse
}

type TaobaoOpenimChatlogsGetResponse struct {
    XMLName xml.Name `xml:"openim_chatlogs_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 聊天记录查询结果
    
    Result   *RoamingMessageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
