package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmessagesendtextAPIRequest 故事机发送文本留言 API请求
// taobao.ailab.aicloud.top.message.sendtext
//
// 故事机文本留言
type TaobaoailabaicloudtopmessagesendtextAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户 id
	_userId string
	// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 留言输入的文本
	_text string
}

// NewTaobaoailabaicloudtopmessagesendtextRequest 初始化TaobaoailabaicloudtopmessagesendtextAPIRequest对象
func NewTaobaoailabaicloudtopmessagesendtextRequest() *TaobaoailabaicloudtopmessagesendtextAPIRequest {
	return &TaobaoailabaicloudtopmessagesendtextAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.sendtext"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopmessagesendtextAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户 id
func (r *TaobaoailabaicloudtopmessagesendtextAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopmessagesendtextAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopmessagesendtextAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetExt() string {
	return r._ext
}

// SetText is Text Setter
// 留言输入的文本
func (r *TaobaoailabaicloudtopmessagesendtextAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaoailabaicloudtopmessagesendtextAPIRequest) GetText() string {
	return r._text
}
