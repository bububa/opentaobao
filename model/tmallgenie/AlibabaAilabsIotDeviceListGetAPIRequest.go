package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsIotDeviceListGetAPIRequest 获取iot设备列表 API请求
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
type AlibabaAilabsIotDeviceListGetAPIRequest struct {
	model.Params
	// 用户open id
	_userOpenId string
	// client id
	_clientId string
}

// NewAlibabaAilabsIotDeviceListGetRequest 初始化AlibabaAilabsIotDeviceListGetAPIRequest对象
func NewAlibabaAilabsIotDeviceListGetRequest() *AlibabaAilabsIotDeviceListGetAPIRequest {
	return &AlibabaAilabsIotDeviceListGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsIotDeviceListGetAPIRequest) Reset() {
	r._userOpenId = ""
	r._clientId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.iot.device.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserOpenId is UserOpenId Setter
// 用户open id
func (r *AlibabaAilabsIotDeviceListGetAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

// SetClientId is ClientId Setter
// client id
func (r *AlibabaAilabsIotDeviceListGetAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsIotDeviceListGetAPIRequest) GetClientId() string {
	return r._clientId
}

var poolAlibabaAilabsIotDeviceListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsIotDeviceListGetRequest()
	},
}

// GetAlibabaAilabsIotDeviceListGetRequest 从 sync.Pool 获取 AlibabaAilabsIotDeviceListGetAPIRequest
func GetAlibabaAilabsIotDeviceListGetAPIRequest() *AlibabaAilabsIotDeviceListGetAPIRequest {
	return poolAlibabaAilabsIotDeviceListGetAPIRequest.Get().(*AlibabaAilabsIotDeviceListGetAPIRequest)
}

// ReleaseAlibabaAilabsIotDeviceListGetAPIRequest 将 AlibabaAilabsIotDeviceListGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsIotDeviceListGetAPIRequest(v *AlibabaAilabsIotDeviceListGetAPIRequest) {
	v.Reset()
	poolAlibabaAilabsIotDeviceListGetAPIRequest.Put(v)
}
