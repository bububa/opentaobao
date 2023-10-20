package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest 获取未读的消息数量 API请求
// taobao.ailab.aicloud.top.message.get.unread.count
//
// 开放获取未读留言数量的接口
type TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// NewTaobaoAilabAicloudTopMessageGetUnreadCountRequest 初始化TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest对象
func NewTaobaoAilabAicloudTopMessageGetUnreadCountRequest() *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest {
	return &TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.get.unread.count"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) GetExt() string {
	return r._ext
}

var poolTaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMessageGetUnreadCountRequest()
	},
}

// GetTaobaoAilabAicloudTopMessageGetUnreadCountRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest
func GetTaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest() *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest {
	return poolTaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest.Get().(*TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest 将 TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest(v *TaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageGetUnreadCountAPIRequest.Put(v)
}
