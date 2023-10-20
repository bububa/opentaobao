package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceSettingsReset 重置设备个性化设置
// taobao.ailab.aicloud.top.device.settings.reset
//
// 重置设备个性化设置
func TaobaoAilabAicloudTopDeviceSettingsReset(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceSettingsResetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceSettingsResetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
