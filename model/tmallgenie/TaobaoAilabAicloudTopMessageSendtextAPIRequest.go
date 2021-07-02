package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageSendtextAPIRequest 故事机发送文本留言 API请求
// taobao.ailab.aicloud.top.message.sendtext
//
// 故事机文本留言
type TaobaoAilabAicloudTopMessageSendtextAPIRequest struct {
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

// NewTaobaoAilabAicloudTopMessageSendtextRequest 初始化TaobaoAilabAicloudTopMessageSendtextAPIRequest对象
func NewTaobaoAilabAicloudTopMessageSendtextRequest() *TaobaoAilabAicloudTopMessageSendtextAPIRequest {
	return &TaobaoAilabAicloudTopMessageSendtextAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.sendtext"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageSendtextAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UserId Setter
// 用户ID，此处传入第三方账户体系的用户 id
func (r *TaobaoAilabAicloudTopMessageSendtextAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetUserId() string {
	return r._userId
}

// Set is UtdId Setter
// 用户设备唯一识别码，长度限制32以内， 建议使用系统接口获取deviceid, 然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageSendtextAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageSendtextAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetExt() string {
	return r._ext
}

// Set is Text Setter
// 留言输入的文本
func (r *TaobaoAilabAicloudTopMessageSendtextAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// Get Text Getter
func (r TaobaoAilabAicloudTopMessageSendtextAPIRequest) GetText() string {
	return r._text
}
