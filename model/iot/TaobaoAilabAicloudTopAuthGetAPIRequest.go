package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopauthgetAPIRequest 登陆 API请求
// taobao.ailab.aicloud.top.auth.get
//
// 登陆
type TaobaoailabaicloudtopauthgetAPIRequest struct {
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

// NewTaobaoailabaicloudtopauthgetRequest 初始化TaobaoailabaicloudtopauthgetAPIRequest对象
func NewTaobaoailabaicloudtopauthgetRequest() *TaobaoailabaicloudtopauthgetAPIRequest {
	return &TaobaoailabaicloudtopauthgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.auth.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopauthgetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopauthgetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopauthgetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetSchema() string {
	return r._schema
}

// SetAppType is AppType Setter
// app类型
func (r *TaobaoailabaicloudtopauthgetAPIRequest) SetAppType(_appType string) error {
	r._appType = _appType
	r.Set("app_type", _appType)
	return nil
}

// GetAppType AppType Getter
func (r TaobaoailabaicloudtopauthgetAPIRequest) GetAppType() string {
	return r._appType
}
