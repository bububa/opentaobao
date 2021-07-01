package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest
零配方案上报设备 API请求
taobao.ailab.aicloud.smarthome.top.genielink.reportdevice

零配方案中设备联网成功之后上报设备 */
type TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest struct {
	model.Params
	// 供应商id
	_vendorId int64
	// 设备id
	_deviceId string
	// 设备状态，online上线，offline下线
	_status string
	// 保留字段json字符串
	_extensions string
}

// New
