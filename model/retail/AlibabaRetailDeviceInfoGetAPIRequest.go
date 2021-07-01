package retail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailDeviceInfoGetAPIRequest
贩卖机设备信息获取 API请求
alibaba.retail.device.info.get

贩卖机设备信息获取 */
type AlibabaRetailDeviceInfoGetAPIRequest struct {
	model.Params
	// 外部设备ID
	_deviceUuid string
	// 阿里设备编码
	_deviceCode string
	// 阿里设备物理ID（32位）
	_deviceSn string
}

// New
