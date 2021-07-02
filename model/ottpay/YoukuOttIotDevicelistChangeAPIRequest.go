package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttIotDevicelistChangeAPIRequest iot设备列表变化接口 API请求
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
type YoukuOttIotDevicelistChangeAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// NewYoukuOttIotDevicelistChangeRequest 初始化YoukuOttIotDevicelistChangeAPIRequest对象
func NewYoukuOttIotDevicelistChangeRequest() *YoukuOttIotDevicelistChangeAPIRequest {
	return &YoukuOttIotDevicelistChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttIotDevicelistChangeAPIRequest) GetApiMethodName() string {
	return "youku.ott.iot.devicelist.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttIotDevicelistChangeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ChangeInfo Setter
// 变更信息
func (r *YoukuOttIotDevicelistChangeAPIRequest) SetChangeInfo(_changeInfo string) error {
	r._changeInfo = _changeInfo
	r.Set("change_info", _changeInfo)
	return nil
}

// Get ChangeInfo Getter
func (r YoukuOttIotDevicelistChangeAPIRequest) GetChangeInfo() string {
	return r._changeInfo
}
