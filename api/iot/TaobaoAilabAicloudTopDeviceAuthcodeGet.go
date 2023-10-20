package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceAuthcodeGet 获取设备授权码
// taobao.ailab.aicloud.top.device.authcode.get
//
// 获取设备授权码
func TaobaoAilabAicloudTopDeviceAuthcodeGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceAuthcodeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
