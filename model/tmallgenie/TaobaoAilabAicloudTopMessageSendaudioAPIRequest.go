package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagesendaudioAPIRequest 发送语音留言 API请求
// taobao.ailab.aicloud.top.message.sendaudio
//
// 将语音的二进制byte[]通过TOP接口发送保存
type TaobaoailabaicloudtopmessagesendaudioAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户 id
	_userId string
	// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 语音的二进制
	_message *model.File
}

// NewTaobaoailabaicloudtopmessagesendaudioRequest 初始化TaobaoailabaicloudtopmessagesendaudioAPIRequest对象
func NewTaobaoailabaicloudtopmessagesendaudioRequest() *TaobaoailabaicloudtopmessagesendaudioAPIRequest {
	return &TaobaoailabaicloudtopmessagesendaudioAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.sendaudio"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopmessagesendaudioAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户 id
func (r *TaobaoailabaicloudtopmessagesendaudioAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopmessagesendaudioAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopmessagesendaudioAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetExt() string {
	return r._ext
}

// SetMessage is Message Setter
// 语音的二进制
func (r *TaobaoailabaicloudtopmessagesendaudioAPIRequest) SetMessage(_message *model.File) error {
	r._message = _message
	r.Set("message", _message)
	return nil
}

// GetMessage Message Getter
func (r TaobaoailabaicloudtopmessagesendaudioAPIRequest) GetMessage() *model.File {
	return r._message
}
