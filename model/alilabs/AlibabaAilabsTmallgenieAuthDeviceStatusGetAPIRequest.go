package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest 设备状态查询 API请求
// alibaba.ailabs.tmallgenie.auth.device.status.get
//
// 提供给天猫精灵定制机厂商 查询设备在线状态值
type AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest struct {
	model.Params
	// 给产品分配的唯一ID（22位）
	_clientId string
	// 精灵用户的唯一外部ID
	_userOpenId string
	// 精灵设备的唯一ID
	_uuid string
}

// NewAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest 初始化AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest() *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) Reset() {
	r._clientId = ""
	r._userOpenId = ""
	r._uuid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.status.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 给产品分配的唯一ID（22位）
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 精灵用户的唯一外部ID
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetUuid is Uuid Setter
// 精灵设备的唯一ID
func (r *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) GetUuid() string {
	return r._uuid
}

var poolAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceStatusGetRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest
func GetAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest() *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest 将 AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest(v *AlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceStatusGetAPIRequest.Put(v)
}
