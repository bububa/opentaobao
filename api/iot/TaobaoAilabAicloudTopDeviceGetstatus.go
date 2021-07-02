package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceGetstatus 获取设备状态
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
func TaobaoAilabAicloudTopDeviceGetstatus(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceGetstatusAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceGetstatusAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceGetstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
