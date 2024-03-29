package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMessageListAPIRequest 获取留言列表 API请求
// taobao.ailab.aicloud.top.message.list
//
// 根据指定参数获取留言列表
type TaobaoAilabAicloudTopMessageListAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 截止时间，yyyy-MM-dd HH:mm:ss.SSS
	_beforeTime string
	// 一次性展示多少条信息
	_limit int64
}

// NewTaobaoAilabAicloudTopMessageListRequest 初始化TaobaoAilabAicloudTopMessageListAPIRequest对象
func NewTaobaoAilabAicloudTopMessageListRequest() *TaobaoAilabAicloudTopMessageListAPIRequest {
	return &TaobaoAilabAicloudTopMessageListAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._beforeTime = ""
	r._limit = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.message.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetExt() string {
	return r._ext
}

// SetBeforeTime is BeforeTime Setter
// 截止时间，yyyy-MM-dd HH:mm:ss.SSS
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) SetBeforeTime(_beforeTime string) error {
	r._beforeTime = _beforeTime
	r.Set("before_time", _beforeTime)
	return nil
}

// GetBeforeTime BeforeTime Getter
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetBeforeTime() string {
	return r._beforeTime
}

// SetLimit is Limit Setter
// 一次性展示多少条信息
func (r *TaobaoAilabAicloudTopMessageListAPIRequest) SetLimit(_limit int64) error {
	r._limit = _limit
	r.Set("limit", _limit)
	return nil
}

// GetLimit Limit Getter
func (r TaobaoAilabAicloudTopMessageListAPIRequest) GetLimit() int64 {
	return r._limit
}

var poolTaobaoAilabAicloudTopMessageListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMessageListRequest()
	},
}

// GetTaobaoAilabAicloudTopMessageListRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMessageListAPIRequest
func GetTaobaoAilabAicloudTopMessageListAPIRequest() *TaobaoAilabAicloudTopMessageListAPIRequest {
	return poolTaobaoAilabAicloudTopMessageListAPIRequest.Get().(*TaobaoAilabAicloudTopMessageListAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMessageListAPIRequest 将 TaobaoAilabAicloudTopMessageListAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMessageListAPIRequest(v *TaobaoAilabAicloudTopMessageListAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMessageListAPIRequest.Put(v)
}
