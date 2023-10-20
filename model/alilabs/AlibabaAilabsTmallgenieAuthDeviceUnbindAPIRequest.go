package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest 解绑设备 API请求
// alibaba.ailabs.tmallgenie.auth.device.unbind
//
// 通过此接口解绑天猫精灵设备
type AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
	// 设备uuid
	_uuid string
}

// NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest 初始化AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest() *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) Reset() {
	r._clientId = ""
	r._userOpenId = ""
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetUuid is Uuid Setter
// 设备uuid
func (r *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) GetUuid() string {
	return r._uuid
}

var poolAlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthDeviceUnbindRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceUnbindRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest
func GetAlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest() *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest 将 AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest(v *AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest.Put(v)
}
