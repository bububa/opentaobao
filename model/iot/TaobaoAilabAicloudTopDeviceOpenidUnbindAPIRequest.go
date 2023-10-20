package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest openTaoBaoId解绑设备 API请求
// taobao.ailab.aicloud.top.device.openid.unbind
//
// openTaoBaoId解绑设备
type TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest struct {
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

// NewTaobaoAilabAicloudTopDeviceOpenidUnbindRequest 初始化TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest对象
func NewTaobaoAilabAicloudTopDeviceOpenidUnbindRequest() *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest {
	return &TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._ext = ""
	r._utdId = ""
	r._openId = ""
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.device.openid.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetUserId() string {
	return r._userId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetExt() string {
	return r._ext
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetOpenId is OpenId Setter
// 淘宝openId
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) GetUuid() string {
	return r._uuid
}

var poolTaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopDeviceOpenidUnbindRequest()
	},
}

// GetTaobaoAilabAicloudTopDeviceOpenidUnbindRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest
func GetTaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest() *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest {
	return poolTaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest.Get().(*TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest 将 TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest(v *TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest.Put(v)
}
