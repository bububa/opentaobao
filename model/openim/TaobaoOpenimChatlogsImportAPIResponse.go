package openim

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
openim单聊消息导入 API返回值 
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能
*/
type TaobaoOpenimChatlogsImportAPIResponse struct {
    model.CommonResponse
    TaobaoOpenimChatlogsImportAPIResponseModel
}

// openim单聊消息导入 成功返回结果
type TaobaoOpenimChatlogsImportAPIResponseModel struct {
    XMLName xml.Name `xml:"openim_chatlogs_import_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    Ret   int64 `json:"ret,omitempty" xml:"ret,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    // 是否成功
    Succ   bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
