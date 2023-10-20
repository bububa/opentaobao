package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagesendAPIRequest 发送留言 API请求
// taobao.ailab.aicloud.top.message.send
//
// 供准入的外部用户实现发送留言功能，APP端发送，设备端读取
type TaobaoailabaicloudtopmessagesendAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 账户体系隔离
	_schema string
	// 用户上传到OSS上的文件路径，不包含域名
	_audioPath string
}

// NewTaobaoailabaicloudtopmessagesendRequest 初始化TaobaoailabaicloudtopmessagesendAPIRequest对象
func NewTaobaoailabaicloudtopmessagesendRequest() *TaobaoailabaicloudtopmessagesendAPIRequest {
	return &TaobaoailabaicloudtopmessagesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopmessagesendAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetExt() string {
	return r._ext
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopmessagesendAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopmessagesendAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetUserId() string {
	return r._userId
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopmessagesendAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetSchema() string {
	return r._schema
}

// SetAudioPath is AudioPath Setter
// 用户上传到OSS上的文件路径，不包含域名
func (r *TaobaoailabaicloudtopmessagesendAPIRequest) SetAudioPath(_audioPath string) error {
	r._audioPath = _audioPath
	r.Set("audio_path", _audioPath)
	return nil
}

// GetAudioPath AudioPath Getter
func (r TaobaoailabaicloudtopmessagesendAPIRequest) GetAudioPath() string {
	return r._audioPath
}
