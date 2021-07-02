package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest 获取设备授权码 API请求
// taobao.ailab.aicloud.top.device.authcode.get
//
// 获取设备授权码
type TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest struct {
	model.Params
	// 账户体系隔离，即硬件接入平台中取得的schema key。
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id，由开发者或厂商自行定义，每一schema key下唯一即可
	_userId string
	// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// NewTaobaoAilabAicloudTopDeviceAuthcodeGetRequest 初始化TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceAuthcodeGetRequest() *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest {
	return &TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.authcode.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// 账户体系隔离，即硬件接入平台中取得的schema key。
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id，由开发者或厂商自行定义，每一schema key下唯一即可
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) GetUserId() string {
	return r._userId
}

// Set is UtdId Setter
// (废弃) 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest) GetExt() string {
	return r._ext
}
