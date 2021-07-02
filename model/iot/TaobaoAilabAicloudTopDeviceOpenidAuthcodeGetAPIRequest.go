package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest 获取openid设备通用授权码 API请求
// taobao.ailab.aicloud.top.device.openid.authcode.get
//
// 获取openid设备通用授权码
type TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest struct {
	model.Params
	// 淘宝openid
	_openId string
	// 账户体系隔离，即硬件接入平台中取得的schema key。
	_schema string
	// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// NewTaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest 初始化TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceOpenidAuthcodeGetRequest() *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest {
	return &TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.openid.authcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenId Setter
// 淘宝openid
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// Get OpenId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) GetOpenId() string {
	return r._openId
}

// Set is Schema Setter
// 账户体系隔离，即硬件接入平台中取得的schema key。
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UtdId Setter
// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest) GetExt() string {
	return r._ext
}
