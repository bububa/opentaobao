package smartstore

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSmartstoreDeviceStatusFeedbackAPIRequest
设备在线状态回流 API请求
taobao.smartstore.device.status.feedback

智能硬件设备状态回流 */
type TaobaoSmartstoreDeviceStatusFeedbackAPIRequest struct {
	model.Params
	// ONLINE_WITH_CONTENT("ONLINE_WITH_CONTENT", "设备在线"), OFFLINE("OFFLINE", "设备断线");
	_status string
	// 设备编码
	_deviceCode string
	// 当前状态的时间
	_statusTime string
}

// NewTaobaoSmartstoreDeviceStatusFeedbackRequest 初始化TaobaoSmartstoreDeviceStatusFeedbackAPIRequest对象
func NewTaobaoSmartstoreDeviceStatusFeedbackRequest() *TaobaoSmartstoreDeviceStatusFeedbackAPIRequest {
	return &TaobaoSmartstoreDeviceStatusFeedbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) GetApiMethodName() string {
	return "taobao.smartstore.device.status.feedback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Status Setter
// ONLINE_WITH_CONTENT("ONLINE_WITH_CONTENT", "设备在线"), OFFLINE("OFFLINE", "设备断线");
func (r *TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) GetStatus() string {
	return r._status
}

// Set is DeviceCode Setter
// 设备编码
func (r *TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// Get DeviceCode Getter
func (r TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// Set is StatusTime Setter
// 当前状态的时间
func (r *TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) SetStatusTime(_statusTime string) error {
	r._statusTime = _statusTime
	r.Set("status_time", _statusTime)
	return nil
}

// Get StatusTime Getter
func (r TaobaoSmartstoreDeviceStatusFeedbackAPIRequest) GetStatusTime() string {
	return r._statusTime
}
