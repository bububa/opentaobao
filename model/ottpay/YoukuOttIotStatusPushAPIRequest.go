package ottpay

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttIotStatusPushAPIRequest iot设备状态变化通知接口 API请求
// youku.ott.iot.status.push
//
// ott iot设备状态通知
type YoukuOttIotStatusPushAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// NewYoukuOttIotStatusPushRequest 初始化YoukuOttIotStatusPushAPIRequest对象
func NewYoukuOttIotStatusPushRequest() *YoukuOttIotStatusPushAPIRequest {
	return &YoukuOttIotStatusPushAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuOttIotStatusPushAPIRequest) Reset() {
	r._changeInfo = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttIotStatusPushAPIRequest) GetApiMethodName() string {
	return "youku.ott.iot.status.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttIotStatusPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttIotStatusPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeInfo is ChangeInfo Setter
// 变更信息
func (r *YoukuOttIotStatusPushAPIRequest) SetChangeInfo(_changeInfo string) error {
	r._changeInfo = _changeInfo
	r.Set("change_info", _changeInfo)
	return nil
}

// GetChangeInfo ChangeInfo Getter
func (r YoukuOttIotStatusPushAPIRequest) GetChangeInfo() string {
	return r._changeInfo
}

var poolYoukuOttIotStatusPushAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuOttIotStatusPushRequest()
	},
}

// GetYoukuOttIotStatusPushRequest 从 sync.Pool 获取 YoukuOttIotStatusPushAPIRequest
func GetYoukuOttIotStatusPushAPIRequest() *YoukuOttIotStatusPushAPIRequest {
	return poolYoukuOttIotStatusPushAPIRequest.Get().(*YoukuOttIotStatusPushAPIRequest)
}

// ReleaseYoukuOttIotStatusPushAPIRequest 将 YoukuOttIotStatusPushAPIRequest 放入 sync.Pool
func ReleaseYoukuOttIotStatusPushAPIRequest(v *YoukuOttIotStatusPushAPIRequest) {
	v.Reset()
	poolYoukuOttIotStatusPushAPIRequest.Put(v)
}
