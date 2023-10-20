package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceUnbindAPIRequest 解绑设备 API请求
// taobao.ailab.aicloud.top.device.unbind
//
// 解绑设备
type TaobaoAilabAicloudTopDeviceUnbindAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
}

// NewTaobaoAilabAicloudTopDeviceUnbindRequest 初始化TaobaoAilabAicloudTopDeviceUnbindAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceUnbindRequest() *TaobaoAilabAicloudTopDeviceUnbindAPIRequest {
	return &TaobaoAilabAicloudTopDeviceUnbindAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceUnbindAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopDeviceUnbindAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceUnbindAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceUnbindAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceUnbindAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopDeviceUnbindAPIRequest) GetExt() string {
	return r._ext
}

var poolTaobaoAilabAicloudTopDeviceUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceUnbindRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceUnbindRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceUnbindAPIRequest
func GetTaobaoAilabAicloudTopDeviceUnbindAPIRequest() *TaobaoAilabAicloudTopDeviceUnbindAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceUnbindAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceUnbindAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceUnbindAPIRequest 将 TaobaoAilabAicloudTopDeviceUnbindAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceUnbindAPIRequest(v *TaobaoAilabAicloudTopDeviceUnbindAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceUnbindAPIRequest.Put(v)
}
