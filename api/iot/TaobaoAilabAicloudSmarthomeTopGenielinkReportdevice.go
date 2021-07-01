package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TaobaoAilabAicloudSmarthomeTopGenielinkReportdevice
零配方案上报设备
taobao.ailab.aicloud.smarthome.top.genielink.reportdevice

零配方案中设备联网成功之后上报设备 */
func TaobaoAilabAicloudSmarthomeTopGenielinkReportdevice(clt *core.SDKClient, req *iot.TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest, session string) (*iot.TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
