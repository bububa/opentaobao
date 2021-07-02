package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceOpenidAuthcodeGet 获取openid设备通用授权码
// taobao.ailab.aicloud.top.device.openid.authcode.get
//
// 获取openid设备通用授权码
func TaobaoAilabAicloudTopDeviceOpenidAuthcodeGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceOpenidAuthcodeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
