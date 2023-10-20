package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceControlCustom 设备控制自定义扩展接口
// taobao.ailab.aicloud.top.device.control.custom
//
// 设备控制自定义扩展接口
func TaobaoAilabAicloudTopDeviceControlCustom(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceControlCustomAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceControlCustomAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceControlCustomAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
