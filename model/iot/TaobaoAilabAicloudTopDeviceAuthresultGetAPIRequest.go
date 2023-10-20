package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest 获取设备授权码验证结果 API请求
// taobao.ailab.aicloud.top.device.authresult.get
//
// 获取设备授权码验证结果
type TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest struct {
	model.Params
	// authCodes信息
	_authCodes []string
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// NewTaobaoAilabAicloudTopDeviceAuthresultGetRequest 初始化TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceAuthresultGetRequest() *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest {
	return &TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) Reset() {
	r._authCodes = r._authCodes[:0]
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.authresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthCodes is AuthCodes Setter
// authCodes信息
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) SetAuthCodes(_authCodes []string) error {
	r._authCodes = _authCodes
	r.Set("auth_codes", _authCodes)
	return nil
}

// GetAuthCodes AuthCodes Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetAuthCodes() []string {
	return r._authCodes
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) GetExt() string {
	return r._ext
}

var poolTaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceAuthresultGetRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceAuthresultGetRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest
func GetTaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest() *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest 将 TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest(v *TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest.Put(v)
}
