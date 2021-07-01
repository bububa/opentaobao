package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuOttIotStatusPushAPIRequest
iot设备状态变化通知接口 API请求
youku.ott.iot.status.push

ott iot设备状态通知 */
type YoukuOttIotStatusPushAPIRequest struct {
	model.Params
	// 变更信息
	_changeInfo string
}

// NewYoukuOttIotStatusPushRequest 初始化YoukuOttIotStatusPushAPIRequest对象
func NewYoukuOttIotStatusPushRequest() *YoukuOttIotStatusPushAPIRequest {
	return &YoukuOttIotStatusPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttIotStatusPushAPIRequest) GetApiMethodName() string {
	return "youku.ott.iot.status.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttIotStatusPushAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ChangeInfo Setter
// 变更信息
func (r *YoukuOttIotStatusPushAPIRequest) SetChangeInfo(_changeInfo string) error {
	r._changeInfo = _changeInfo
	r.Set("change_info", _changeInfo)
	return nil
}

// Get ChangeInfo Getter
func (r YoukuOttIotStatusPushAPIRequest) GetChangeInfo() string {
	return r._changeInfo
}
