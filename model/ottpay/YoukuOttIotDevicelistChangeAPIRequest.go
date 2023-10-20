package ottpay

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttIotDevicelistChangeAPIRequest) Reset() {
	r._changeInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttIotDevicelistChangeAPIRequest) GetApiMethodName() string {
	return "youku.ott.iot.devicelist.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttIotDevicelistChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttIotDevicelistChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeInfo is ChangeInfo Setter
// 变更信息
func (r *YoukuOttIotDevicelistChangeAPIRequest) SetChangeInfo(_changeInfo string) error {
	r._changeInfo = _changeInfo
	r.Set("change_info", _changeInfo)
	return nil
}

// GetChangeInfo ChangeInfo Getter
func (r YoukuOttIotDevicelistChangeAPIRequest) GetChangeInfo() string {
	return r._changeInfo
}

var poolYoukuOttIotDevicelistChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttIotDevicelistChangeRequest()
	},
}

// GetYoukuOttIotDevicelistChangeRequest 从 sync.Pool 获取 YoukuOttIotDevicelistChangeAPIRequest
func GetYoukuOttIotDevicelistChangeAPIRequest() *YoukuOttIotDevicelistChangeAPIRequest {
	return poolYoukuOttIotDevicelistChangeAPIRequest.Get().(*YoukuOttIotDevicelistChangeAPIRequest)
}

// ReleaseYoukuOttIotDevicelistChangeAPIRequest 将 YoukuOttIotDevicelistChangeAPIRequest 放入 sync.Pool
func ReleaseYoukuOttIotDevicelistChangeAPIRequest(v *YoukuOttIotDevicelistChangeAPIRequest) {
	v.Reset()
	poolYoukuOttIotDevicelistChangeAPIRequest.Put(v)
}
