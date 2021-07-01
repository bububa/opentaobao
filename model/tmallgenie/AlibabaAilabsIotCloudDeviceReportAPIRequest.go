package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAilabsIotCloudDeviceReportAPIRequest
天猫精灵云云接入设备状态、事件上报接口 API请求
alibaba.ailabs.iot.cloud.device.report

承接天猫精灵云云接入设备的状态、事件上报 */
type AlibabaAilabsIotCloudDeviceReportAPIRequest struct {
	model.Params
	// 上报总入参
	_cloudReportParam *CloudReportParam
}

// New
