package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlPauseandresume 设备播放暂停
// taobao.ailab.aicloud.top.device.control.pauseandresume
//
// 设备播放暂停
func TaobaoAilabAicloudTopDeviceControlPauseandresume(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlPauseandresumeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
