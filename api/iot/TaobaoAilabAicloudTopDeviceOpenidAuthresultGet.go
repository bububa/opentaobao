package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceOpenidAuthresultGet 获取openId设备授权码验证结果
// taobao.ailab.aicloud.top.device.openid.authresult.get
//
// 获取openId设备授权码验证结果
func TaobaoAilabAicloudTopDeviceOpenidAuthresultGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceOpenidAuthresultGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
