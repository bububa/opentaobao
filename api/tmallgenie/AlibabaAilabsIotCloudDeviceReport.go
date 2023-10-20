package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsIotCloudDeviceReport 天猫精灵云云接入设备状态、事件上报接口
// alibaba.ailabs.iot.cloud.device.report
//
// 承接天猫精灵云云接入设备的状态、事件上报
func AlibabaAilabsIotCloudDeviceReport(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotCloudDeviceReportAPIRequest, resp *tmallgenie.AlibabaAilabsIotCloudDeviceReportAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
