package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTvpubadminDeviceStatsAPIRequest
获取设备统计数据 API请求
yunos.tvpubadmin.device.stats

获取设备统计数据 */
type YunosTvpubadminDeviceStatsAPIRequest struct {
	model.Params
	// 厂商名称
	_factoryName string
	// 设备型号
	_deviceModel string
}

// New
