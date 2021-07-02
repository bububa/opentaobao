package aligenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieDeviceUnbindAPIRequest 设备解绑操作接口 API请求
// alibaba.ailabs.aligenie.device.unbind
//
// 让开发者能根据设备ID进行解绑操作的接口
type AlibabaAilabsAligenieDeviceUnbindAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 账户体系隔离字符串
	_schema string
	// 欲解绑的设备ID
	_uuid string
}

// NewAlibabaAilabsAligenieDeviceUnbindRequest 初始化AlibabaAilabsAligenieDeviceUnbindAPIRequest对象
func NewAlibabaAilabsAligenieDeviceUnbindRequest() *AlibabaAilabsAligenieDeviceUnbindAPIRequest {
	return &AlibabaAilabsAligenieDeviceUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieDeviceUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieDeviceUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaAilabsAligenieDeviceUnbindAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r AlibabaAilabsAligenieDeviceUnbindAPIRequest) GetExt() string {
	return r._ext
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaAilabsAligenieDeviceUnbindAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAilabsAligenieDeviceUnbindAPIRequest) GetUserId() string {
	return r._userId
}

// SetSchema is Schema Setter
// 账户体系隔离字符串
func (r *AlibabaAilabsAligenieDeviceUnbindAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaAilabsAligenieDeviceUnbindAPIRequest) GetSchema() string {
	return r._schema
}

// SetUuid is Uuid Setter
// 欲解绑的设备ID
func (r *AlibabaAilabsAligenieDeviceUnbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAilabsAligenieDeviceUnbindAPIRequest) GetUuid() string {
	return r._uuid
}
