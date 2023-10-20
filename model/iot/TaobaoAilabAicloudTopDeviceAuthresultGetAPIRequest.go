package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceauthresultgetAPIRequest 获取设备授权码验证结果 API请求
// taobao.ailab.aicloud.top.device.authresult.get
//
// 获取设备授权码验证结果
type TaobaoailabaicloudtopdeviceauthresultgetAPIRequest struct {
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

// NewTaobaoailabaicloudtopdeviceauthresultgetRequest 初始化TaobaoailabaicloudtopdeviceauthresultgetAPIRequest对象
func NewTaobaoailabaicloudtopdeviceauthresultgetRequest() *TaobaoailabaicloudtopdeviceauthresultgetAPIRequest {
	return &TaobaoailabaicloudtopdeviceauthresultgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.authresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthCodes is AuthCodes Setter
// authCodes信息
func (r *TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) SetAuthCodes(_authCodes []string) error {
	r._authCodes = _authCodes
	r.Set("auth_codes", _authCodes)
	return nil
}

// GetAuthCodes AuthCodes Getter
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetAuthCodes() []string {
	return r._authCodes
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopdeviceauthresultgetAPIRequest) GetExt() string {
	return r._ext
}
