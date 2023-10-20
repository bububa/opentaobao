package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartstoredevicestatusfeedbackAPIRequest 设备在线状态回流 API请求
// taobao.smartstore.device.status.feedback
//
// 智能硬件设备状态回流
type TaobaosmartstoredevicestatusfeedbackAPIRequest struct {
	model.Params
	// ONLINE_WITH_CONTENT("ONLINE_WITH_CONTENT", "设备在线"), OFFLINE("OFFLINE", "设备断线");
	_status string
	// 设备编码
	_deviceCode string
	// 当前状态的时间
	_statusTime string
}

// NewTaobaosmartstoredevicestatusfeedbackRequest 初始化TaobaosmartstoredevicestatusfeedbackAPIRequest对象
func NewTaobaosmartstoredevicestatusfeedbackRequest() *TaobaosmartstoredevicestatusfeedbackAPIRequest {
	return &TaobaosmartstoredevicestatusfeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosmartstoredevicestatusfeedbackAPIRequest) GetApiMethodName() string {
	return "taobao.smartstore.device.status.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosmartstoredevicestatusfeedbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosmartstoredevicestatusfeedbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// ONLINE_WITH_CONTENT(&#34;ONLINE_WITH_CONTENT&#34;, &#34;设备在线&#34;), OFFLINE(&#34;OFFLINE&#34;, &#34;设备断线&#34;);
func (r *TaobaosmartstoredevicestatusfeedbackAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaosmartstoredevicestatusfeedbackAPIRequest) GetStatus() string {
	return r._status
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *TaobaosmartstoredevicestatusfeedbackAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r TaobaosmartstoredevicestatusfeedbackAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetStatusTime is StatusTime Setter
// 当前状态的时间
func (r *TaobaosmartstoredevicestatusfeedbackAPIRequest) SetStatusTime(_statusTime string) error {
	r._statusTime = _statusTime
	r.Set("status_time", _statusTime)
	return nil
}

// GetStatusTime StatusTime Getter
func (r TaobaosmartstoredevicestatusfeedbackAPIRequest) GetStatusTime() string {
	return r._statusTime
}
