package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceRoadStatusResetAPIRequest
贩卖机货道解锁 API请求
alibaba.retail.device.road.status.reset

贩卖机货道解锁 */
type AlibabaRetailDeviceRoadStatusResetAPIRequest struct {
	model.Params
	// 设备外部编码
	_deviceUuid string
	// 阿里设备编码
	_deviceCode string
	// 阿里设备物理编码
	_deviceSn string
	// 货道编码
	_roadNoList []string
}

// New
