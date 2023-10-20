package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopfeedlistdeleteAPIRequest 删除单条对话流信息 API请求
// taobao.ailab.aicloud.top.feedlist.delete
//
// 删除指定的某一条对话流信息
type TaobaoailabaicloudtopfeedlistdeleteAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 消息的唯一标识
	_sentenceId string
}

// NewTaobaoailabaicloudtopfeedlistdeleteRequest 初始化TaobaoailabaicloudtopfeedlistdeleteAPIRequest对象
func NewTaobaoailabaicloudtopfeedlistdeleteRequest() *TaobaoailabaicloudtopfeedlistdeleteAPIRequest {
	return &TaobaoailabaicloudtopfeedlistdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.feedlist.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopfeedlistdeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopfeedlistdeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopfeedlistdeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopfeedlistdeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetSentenceId is SentenceId Setter
// 消息的唯一标识
func (r *TaobaoailabaicloudtopfeedlistdeleteAPIRequest) SetSentenceId(_sentenceId string) error {
	r._sentenceId = _sentenceId
	r.Set("sentence_id", _sentenceId)
	return nil
}

// GetSentenceId SentenceId Getter
func (r TaobaoailabaicloudtopfeedlistdeleteAPIRequest) GetSentenceId() string {
	return r._sentenceId
}
