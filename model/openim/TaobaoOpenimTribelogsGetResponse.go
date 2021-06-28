package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openim 群聊天记录导出接口 APIResponse
taobao.openim.tribelogs.get

获取openim账号的群聊天记录
*/
type TaobaoOpenimTribelogsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"openim_tribelogs_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    RetCode   int64 `json:"retCode,omitempty" xml:"