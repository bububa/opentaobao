package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceAuthresultGet 获取设备授权码验证结果
// taobao.ailab.aicloud.top.device.authresult.get
//
// 获取设备授权码验证结果
func TaobaoAilabAicloudTopDeviceAuthresultGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceAuthresultGetAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceAuthresultGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
