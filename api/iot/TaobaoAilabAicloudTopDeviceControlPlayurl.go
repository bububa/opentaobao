package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlPlayurl 点播url
// taobao.ailab.aicloud.top.device.control.playurl
//
// 点播url
func TaobaoAilabAicloudTopDeviceControlPlayurl(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlPlayurlAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceControlPlayurlAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
