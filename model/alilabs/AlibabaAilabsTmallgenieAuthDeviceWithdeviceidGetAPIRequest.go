package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest 根据三方ID查询设备注册激活信息 API请求
// alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get
//
// 根据三方ID查询设备注册激活信息
type AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest struct {
	model.Params
	// 设备产品ID
	_clientId string
	// mac地址
	_mac string
}

// NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest 初始化AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) Reset() {
	r._clientId = ""
	r._mac = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 设备产品ID
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetClientId() string {
	return r._clientId
}

// SetMac is Mac Setter
// mac地址
func (r *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) GetMac() string {
	return r._mac
}

var poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest
func GetAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest() *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest 将 AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest(v *AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest.Put(v)
}
