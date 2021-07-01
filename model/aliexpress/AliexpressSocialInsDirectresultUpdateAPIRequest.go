package aliexpress

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV更新INS私信发送的结果 API请求
aliexpress.social.ins.directresult.update

ISV更新INS私信发送的结果
*/
type AliexpressSocialInsDirectresultUpdateAPIRequest struct {
    model.Params
    // 回调id,在获取图片时会返回
    _id   int64
    // 接受消息人INS_ID，也就是查询图片时的request_ins_id
    _receiveInsId   string
    // ISV发送私信人的INS_ID
    _senderInsId   string
    // 1.成功，2.失败。
    _result   int64
}

// 初始化AliexpressSocialInsDirectresultUpdateAPIRequest对象
func NewAliexpressSocialInsDirectresultUpdateRequest() *AliexpressSocialInsDirectresultUpdateAPIRequest{
    return &AliexpressSocialInsDirectresultUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetApiMethodName() string {
    return "aliexpress.social.ins.directresult.update"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Id Setter
// 回调id,在获取图片时会返回
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetId(_id int64) error {
    r._id = _id
    r.Set("id", _id)
    return nil
}

// Id Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetId() int64 {
    return r._id
}
// ReceiveInsId Setter
// 接受消息人INS_ID，也就是查询图片时的request_ins_id
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetReceiveInsId(_receiveInsId string) error {
    r._receiveInsId = _receiveInsId
    r.Set("receive_ins_id", _receiveInsId)
    return nil
}

// ReceiveInsId Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetReceiveInsId() string {
    return r._receiveInsId
}
// SenderInsId Setter
// ISV发送私信人的INS_ID
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetSenderInsId(_senderInsId string) error {
    r._senderInsId = _senderInsId
    r.Set("sender_ins_id", _senderInsId)
    return nil
}

// SenderInsId Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetSenderInsId() string {
    return r._senderInsId
}
// Result Setter
// 1.成功，2.失败。
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetResult(_result int64) error {
    r._result = _result
    r.Set("result", _result)
    return nil
}

// Result Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetResult() int64 {
    return r._result
}
