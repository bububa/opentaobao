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

// New
