package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest 获取openid设备通用授权码 API请求
// taobao.ailab.aicloud.top.device.openid.authcode.get
//
// 获取openid设备通用授权码
type TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest struct {
	model.Params
	// 账户体系隔离，即硬件接入平台中取得的schema key。
	_schema string
	// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 淘宝openid
	_openId string
}

// NewTaobaoailabaicloudtopdeviceopenidauthcodegetRequest 初始化TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest对象
func NewTaobaoailabaicloudtopdeviceopenidauthcodegetRequest() *TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest {
	return &TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.openid.authcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离，即硬件接入平台中取得的schema key。
func (r *TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUtdId is UtdId Setter
// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetExt() string {
	return r._ext
}

// SetOpenId is OpenId Setter
// 淘宝openid
func (r *TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest) GetOpenId() string {
	return r._openId
}
