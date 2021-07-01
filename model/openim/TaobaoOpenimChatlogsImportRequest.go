package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim单聊消息导入 API请求
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能
*/
type TaobaoOpenimChatlogsImportAPIRequest struct {
    model.Params
    // 消息序列
    _messages   []TextMessage
}

// 初始化TaobaoOpenimChatlogsImportAPIRequest对象
func NewTaobaoOpenimChatlogsImportRequest() *TaobaoOpenimChatlogsImportAPIRequest{
    return &TaobaoOpenimChatlogsImportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimChatlogsImportAPIRequest) GetApiMethodName() string {
    return "taobao.openim.chatlogs.import"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimChatlogsImportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Messages Setter
// 消息序列
func (r *TaobaoOpenimChatlogsImportAPIRequest) SetMessages(_messages []TextMessage) error {
    r._messages = _messages
    r.Set("messages", _messages)
    return nil
}

// Messages Getter
func (r TaobaoOpenimChatlogsImportAPIRequest) GetMessages() []TextMessage {
    return r._messages
}
