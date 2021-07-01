package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMcDeviceCircleCheckAPIRequest
云码设备圈选情况查询 API请求
tmall.mc.device.circle.check

云码设备圈选情况查询 */
type TmallMcDeviceCircleCheckAPIRequest struct {
	model.Params
	// 外部设备编码
	_outerCode string
	// 渠道编码
	_channelId string
}

// New
