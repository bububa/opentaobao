package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openim群聊天记录导入 API请求
taobao.openim.tribelogs.import

openim群聊天记录导入
*/
type TaobaoOpenimTribelogsImportAPIRequest struct {
    model.Params
    // 群号。必须为已存在的群，且群主属于本app
    _tribeId   int64
    // 消息列表
    _messages   []TribeTextMessage
}

// 初始化TaobaoOpenimTribelogsImportAPIRequest对象
func NewTaobaoOpenimTribelogsImportRequest() *TaobaoOpenimTribelogsImportAPIRequest{
    return &TaobaoOpenimTribelogsImportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimTribelogsImportAPIRequest) GetApiMethodName() string {
    return "taobao.openim.tribelogs.import"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimTribelogsImportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TribeId Setter
// 群号。必须为已存在的群，且群主属于本app
func (r *TaobaoOpenimTribelogsImportAPIRequest) SetTribeId(_tribeId int64) error {
    r._tribeId = _tribeId
    r.Set("tribe_id", _tribeId)
    return nil
}

// TribeId Getter
func (r TaobaoOpenimTribelogsImportAPIRequest) GetTribeId() int64 {
    return r._tribeId
}
// Messages Setter
// 消息列表
func (r *TaobaoOpenimTribelogsImportAPIRequest) SetMessages(_messages []TribeTextMessage) error {
    r._messages = _messages
    r.Set("messages", _messages)
    return nil
}

// Messages Getter
func (r TaobaoOpenimTribelogsImportAPIRequest) GetMessages() []TribeTextMessage {
    return r._messages
}
