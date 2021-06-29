package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openim群聊天记录导入 APIResponse
taobao.openim.tribelogs.import

openim群聊天记录导入
*/
type TaobaoOpenimTribelogsImportAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimTribelogsImportResponse
}

type TaobaoOpenimTribelogsImportResponse struct {
    XMLName xml.Name `xml:"openim_tribelogs_import_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    Ret   int64 `json:"ret,omitempty" xml:"ret,omitempty"`

    
    // 错误信息
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`

    
    // 是否成功
    
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`

    
}
