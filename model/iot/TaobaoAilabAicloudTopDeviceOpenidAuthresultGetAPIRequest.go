package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest 获取openId设备授权码验证结果 API请求
// taobao.ailab.aicloud.top.device.openid.authresult.get
//
// 获取openId设备授权码验证结果
type TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest struct {
	model.Params
	// authcode list
	_authCodes []string
	// 账户体系隔离
	_schema string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 淘宝openid
	_openId string
}

// NewTaobaoailabaicloudtopdeviceopenidauthresultgetRequest 初始化TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest对象
func NewTaobaoailabaicloudtopdeviceopenidauthresultgetRequest() *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest {
	return &TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.openid.authresult.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthCodes is AuthCodes Setter
// authcode list
func (r *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) SetAuthCodes(_authCodes []string) error {
	r._authCodes = _authCodes
	r.Set("auth_codes", _authCodes)
	return nil
}

// GetAuthCodes AuthCodes Getter
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetAuthCodes() []string {
	return r._authCodes
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetUserId() string {
	return r._userId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetExt() string {
	return r._ext
}

// SetOpenId is OpenId Setter
// 淘宝openid
func (r *TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest) GetOpenId() string {
	return r._openId
}
