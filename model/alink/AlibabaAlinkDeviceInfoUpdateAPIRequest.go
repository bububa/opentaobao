package alink

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlinkDeviceInfoUpdateAPIRequest 更新设备昵称等信息 API请求
// alibaba.alink.device.info.update
//
// 更新设备昵称等信息
type AlibabaAlinkDeviceInfoUpdateAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 设备昵称
	_nickName string
}

// NewAlibabaAlinkDeviceInfoUpdateRequest 初始化AlibabaAlinkDeviceInfoUpdateAPIRequest对象
func NewAlibabaAlinkDeviceInfoUpdateRequest() *AlibabaAlinkDeviceInfoUpdateAPIRequest {
	return &AlibabaAlinkDeviceInfoUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlinkDeviceInfoUpdateAPIRequest) Reset() {
	r._uuid = ""
	r._nickName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.info.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceInfoUpdateAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetUuid() string {
	return r._uuid
}

// SetNickName is NickName Setter
// 设备昵称
func (r *AlibabaAlinkDeviceInfoUpdateAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// GetNickName NickName Getter
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetNickName() string {
	return r._nickName
}

var poolAlibabaAlinkDeviceInfoUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlinkDeviceInfoUpdateRequest()
	},
}

// GetAlibabaAlinkDeviceInfoUpdateRequest 从 sync.Pool 获取 AlibabaAlinkDeviceInfoUpdateAPIRequest
func GetAlibabaAlinkDeviceInfoUpdateAPIRequest() *AlibabaAlinkDeviceInfoUpdateAPIRequest {
	return poolAlibabaAlinkDeviceInfoUpdateAPIRequest.Get().(*AlibabaAlinkDeviceInfoUpdateAPIRequest)
}

// ReleaseAlibabaAlinkDeviceInfoUpdateAPIRequest 将 AlibabaAlinkDeviceInfoUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlinkDeviceInfoUpdateAPIRequest(v *AlibabaAlinkDeviceInfoUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlinkDeviceInfoUpdateAPIRequest.Put(v)
}
