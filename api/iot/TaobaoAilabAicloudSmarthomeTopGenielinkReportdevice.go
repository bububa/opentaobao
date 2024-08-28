package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudSmarthomeTopGenielinkReportdevice 零配方案上报设备
// taobao.ailab.aicloud.smarthome.top.genielink.reportdevice
//
// 零配方案中设备联网成功之后上报设备
func TaobaoAilabAicloudSmarthomeTopGenielinkReportdevice(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIRequest, resp *iot.TaobaoAilabAicloudSmarthomeTopGenielinkReportdeviceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
