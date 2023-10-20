package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopdevicesettingsresetAPIRequest 重置设备个性化设置 API请求
// taobao.ailab.aicloud.top.device.settings.reset
//
// 重置设备个性化设置
type TaobaoailabaicloudtopdevicesettingsresetAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 设备id
	_param1 string
}

// NewTaobaoailabaicloudtopdevicesettingsresetRequest 初始化TaobaoailabaicloudtopdevicesettingsresetAPIRequest对象
func NewTaobaoailabaicloudtopdevicesettingsresetRequest() *TaobaoailabaicloudtopdevicesettingsresetAPIRequest {
	return &TaobaoailabaicloudtopdevicesettingsresetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.settings.reset"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtopdevicesettingsresetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtopdevicesettingsresetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtopdevicesettingsresetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtopdevicesettingsresetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetExt() string {
	return r._ext
}

// SetParam1 is Param1 Setter
// 设备id
func (r *TaobaoailabaicloudtopdevicesettingsresetAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtopdevicesettingsresetAPIRequest) GetParam1() string {
	return r._param1
}
