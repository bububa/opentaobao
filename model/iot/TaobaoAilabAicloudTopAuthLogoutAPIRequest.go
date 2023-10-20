package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopauthlogoutAPIRequest 登出 API请求
// taobao.ailab.aicloud.top.auth.logout
//
// 登出
type TaobaoailabaicloudtopauthlogoutAPIRequest struct {
	model.Params
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 账户体系隔离，建议传入设备uuid
	_schema string
}

// NewTaobaoailabaicloudtopauthlogoutRequest 初始化TaobaoailabaicloudtopauthlogoutAPIRequest对象
func NewTaobaoailabaicloudtopauthlogoutRequest() *TaobaoailabaicloudtopauthlogoutAPIRequest {
	return &TaobaoailabaicloudtopauthlogoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopauthlogoutAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.auth.logout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopauthlogoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopauthlogoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopauthlogoutAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopauthlogoutAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopauthlogoutAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopauthlogoutAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetSchema is Schema Setter
// 账户体系隔离，建议传入设备uuid
func (r *TaobaoailabaicloudtopauthlogoutAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopauthlogoutAPIRequest) GetSchema() string {
	return r._schema
}
