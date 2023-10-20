package aliexpress

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialinsdirectresultupdateAPIRequest ISV更新INS私信发送的结果 API请求
// aliexpress.social.ins.directresult.update
//
// ISV更新INS私信发送的结果
type AliexpresssocialinsdirectresultupdateAPIRequest struct {
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

// NewAliexpresssocialinsdirectresultupdateRequest 初始化AliexpresssocialinsdirectresultupdateAPIRequest对象
func NewAliexpresssocialinsdirectresultupdateRequest() *AliexpresssocialinsdirectresultupdateAPIRequest {
	return &AliexpresssocialinsdirectresultupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.ins.directresult.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiveInsId is ReceiveInsId Setter
// 接受消息人INS_ID，也就是查询图片时的request_ins_id
func (r *AliexpresssocialinsdirectresultupdateAPIRequest) SetReceiveInsId(_receiveInsId string) error {
	r._receiveInsId = _receiveInsId
	r.Set("receive_ins_id", _receiveInsId)
	return nil
}

// GetReceiveInsId ReceiveInsId Getter
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetReceiveInsId() string {
	return r._receiveInsId
}

// SetSenderInsId is SenderInsId Setter
// ISV发送私信人的INS_ID
func (r *AliexpresssocialinsdirectresultupdateAPIRequest) SetSenderInsId(_senderInsId string) error {
	r._senderInsId = _senderInsId
	r.Set("sender_ins_id", _senderInsId)
	return nil
}

// GetSenderInsId SenderInsId Getter
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetSenderInsId() string {
	return r._senderInsId
}

// SetId is Id Setter
// 回调id,在获取图片时会返回
func (r *AliexpresssocialinsdirectresultupdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetId() int64 {
	return r._id
}

// SetResult is Result Setter
// 1.成功，2.失败。
func (r *AliexpresssocialinsdirectresultupdateAPIRequest) SetResult(_result int64) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r AliexpresssocialinsdirectresultupdateAPIRequest) GetResult() int64 {
	return r._result
}
