package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceopenidunbindAPIRequest openTaoBaoId解绑设备 API请求
// taobao.ailab.aicloud.top.device.openid.unbind
//
// openTaoBaoId解绑设备
type TaobaoailabaicloudtopdeviceopenidunbindAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 淘宝openId
	_openId string
	// 设备uuid
	_uuid string
}

// NewTaobaoailabaicloudtopdeviceopenidunbindRequest 初始化TaobaoailabaicloudtopdeviceopenidunbindAPIRequest对象
func NewTaobaoailabaicloudtopdeviceopenidunbindRequest() *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest {
	return &TaobaoailabaicloudtopdeviceopenidunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.openid.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetUserId() string {
	return r._userId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetExt() string {
	return r._ext
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetOpenId is OpenId Setter
// 淘宝openId
func (r *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoailabaicloudtopdeviceopenidunbindAPIRequest) GetUuid() string {
	return r._uuid
}
