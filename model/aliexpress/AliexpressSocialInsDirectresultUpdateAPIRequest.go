package aliexpress

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialInsDirectresultUpdateAPIRequest ISV更新INS私信发送的结果 API请求
// aliexpress.social.ins.directresult.update
//
// ISV更新INS私信发送的结果
type AliexpressSocialInsDirectresultUpdateAPIRequest struct {
	model.Params
	// 接受消息人INS_ID，也就是查询图片时的request_ins_id
	_receiveInsId string
	// ISV发送私信人的INS_ID
	_senderInsId string
	// 回调id,在获取图片时会返回
	_id int64
	// 1.成功，2.失败。
	_result int64
}

// NewAliexpressSocialInsDirectresultUpdateRequest 初始化AliexpressSocialInsDirectresultUpdateAPIRequest对象
func NewAliexpressSocialInsDirectresultUpdateRequest() *AliexpressSocialInsDirectresultUpdateAPIRequest {
	return &AliexpressSocialInsDirectresultUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) Reset() {
	r._receiveInsId = ""
	r._senderInsId = ""
	r._id = 0
	r._result = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.ins.directresult.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiveInsId is ReceiveInsId Setter
// 接受消息人INS_ID，也就是查询图片时的request_ins_id
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetReceiveInsId(_receiveInsId string) error {
	r._receiveInsId = _receiveInsId
	r.Set("receive_ins_id", _receiveInsId)
	return nil
}

// GetReceiveInsId ReceiveInsId Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetReceiveInsId() string {
	return r._receiveInsId
}

// SetSenderInsId is SenderInsId Setter
// ISV发送私信人的INS_ID
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetSenderInsId(_senderInsId string) error {
	r._senderInsId = _senderInsId
	r.Set("sender_ins_id", _senderInsId)
	return nil
}

// GetSenderInsId SenderInsId Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetSenderInsId() string {
	return r._senderInsId
}

// SetId is Id Setter
// 回调id,在获取图片时会返回
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetId() int64 {
	return r._id
}

// SetResult is Result Setter
// 1.成功，2.失败。
func (r *AliexpressSocialInsDirectresultUpdateAPIRequest) SetResult(_result int64) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r AliexpressSocialInsDirectresultUpdateAPIRequest) GetResult() int64 {
	return r._result
}

var poolAliexpressSocialInsDirectresultUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressSocialInsDirectresultUpdateRequest()
	},
}

// GetAliexpressSocialInsDirectresultUpdateRequest 从 sync.Pool 获取 AliexpressSocialInsDirectresultUpdateAPIRequest
func GetAliexpressSocialInsDirectresultUpdateAPIRequest() *AliexpressSocialInsDirectresultUpdateAPIRequest {
	return poolAliexpressSocialInsDirectresultUpdateAPIRequest.Get().(*AliexpressSocialInsDirectresultUpdateAPIRequest)
}

// ReleaseAliexpressSocialInsDirectresultUpdateAPIRequest 将 AliexpressSocialInsDirectresultUpdateAPIRequest 放入 sync.Pool
func ReleaseAliexpressSocialInsDirectresultUpdateAPIRequest(v *AliexpressSocialInsDirectresultUpdateAPIRequest) {
	v.Reset()
	poolAliexpressSocialInsDirectresultUpdateAPIRequest.Put(v)
}
