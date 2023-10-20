package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceListAPIRequest 获取用户设备列表 API请求
// alibaba.ailabs.tmallgenie.auth.device.list
//
// 通过此接口获取用户绑定的设备信息列表
type AlibabaAilabsTmallgenieAuthDeviceListAPIRequest struct {
	model.Params
	// 客户id
	_clientId string
	// 用户开放id
	_userOpenId string
}

// NewAlibabaAilabsTmallgenieAuthDeviceListRequest 初始化AlibabaAilabsTmallgenieAuthDeviceListAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceListRequest() *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceListAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) Reset() {
	r._clientId = ""
	r._userOpenId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetClientId is ClientId Setter
// 客户id
func (r *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) SetClientId(_clientId string) error {
	r._clientId = _clientId
	r.Set("client_id", _clientId)
	return nil
}

// GetClientId ClientId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetClientId() string {
	return r._clientId
}

// SetUserOpenId is UserOpenId Setter
// 用户开放id
func (r *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) SetUserOpenId(_userOpenId string) error {
	r._userOpenId = _userOpenId
	r.Set("user_open_id", _userOpenId)
	return nil
}

// GetUserOpenId UserOpenId Getter
func (r AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) GetUserOpenId() string {
	return r._userOpenId
}

var poolAlibabaAilabsTmallgenieAuthDeviceListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthDeviceListRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceListRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceListAPIRequest
func GetAlibabaAilabsTmallgenieAuthDeviceListAPIRequest() *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthDeviceListAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthDeviceListAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceListAPIRequest 将 AlibabaAilabsTmallgenieAuthDeviceListAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceListAPIRequest(v *AlibabaAilabsTmallgenieAuthDeviceListAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceListAPIRequest.Put(v)
}
