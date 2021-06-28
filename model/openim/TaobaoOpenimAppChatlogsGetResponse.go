package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openim应用聊天记录查询 APIResponse
taobao.openim.app.chatlogs.get

查询openim应用的聊天记录
*/
type TaobaoOpenimAppChatlogsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_app_chatlogs_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 查询结果
    
    Result   *EsMessageResult `json:"result,omitempty" xml:"