package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopAuthGetAPIRequest 登陆 API请求
// taobao.ailab.aicloud.top.auth.get
//
// 登陆
type TaobaoAilabAicloudTopAuthGetAPIRequest struct {
	model.Params
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 账户体系隔离
	_schema string
	// app类型
	_appType string
}

// NewTaobaoAilabAicloudTopAuthGetRequest 初始化TaobaoAilabAicloudTopAuthGetAPIRequest对象
func NewTaobaoAilabAicloudTopAuthGetRequest() *TaobaoAilabAicloudTopAuthGetAPIRequest {
	return &TaobaoAilabAicloudTopAuthGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopAuthGetAPIRequest) Reset() {
	r._userId = ""
	r._utdId = ""
	r._schema = ""
	r._appType = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.auth.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopAuthGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopAuthGetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopAuthGetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetSchema() string {
	return r._schema
}

// SetAppType is AppType Setter
// app类型
func (r *TaobaoAilabAicloudTopAuthGetAPIRequest) SetAppType(_appType string) error {
	r._appType = _appType
	r.Set("app_type", _appType)
	return nil
}

// GetAppType AppType Getter
func (r TaobaoAilabAicloudTopAuthGetAPIRequest) GetAppType() string {
	return r._appType
}

var poolTaobaoAilabAicloudTopAuthGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopAuthGetRequest()
	},
}

// GetTaobaoAilabAicloudTopAuthGetRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopAuthGetAPIRequest
func GetTaobaoAilabAicloudTopAuthGetAPIRequest() *TaobaoAilabAicloudTopAuthGetAPIRequest {
	return poolTaobaoAilabAicloudTopAuthGetAPIRequest.Get().(*TaobaoAilabAicloudTopAuthGetAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopAuthGetAPIRequest 将 TaobaoAilabAicloudTopAuthGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopAuthGetAPIRequest(v *TaobaoAilabAicloudTopAuthGetAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopAuthGetAPIRequest.Put(v)
}
