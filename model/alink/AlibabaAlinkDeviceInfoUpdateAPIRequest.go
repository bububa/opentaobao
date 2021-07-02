package alink

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alink.device.info.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Uuid Setter
// 设备id
func (r *AlibabaAlinkDeviceInfoUpdateAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// Get Uuid Getter
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetUuid() string {
	return r._uuid
}

// Set is NickName Setter
// 设备昵称
func (r *AlibabaAlinkDeviceInfoUpdateAPIRequest) SetNickName(_nickName string) error {
	r._nickName = _nickName
	r.Set("nick_name", _nickName)
	return nil
}

// Get NickName Getter
func (r AlibabaAlinkDeviceInfoUpdateAPIRequest) GetNickName() string {
	return r._nickName
}
