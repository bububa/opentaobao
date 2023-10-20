package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceDetailinfoGet 获取设备详细信息
// taobao.ailab.aicloud.top.device.detailinfo.get
//
// 获取设备详细信息
func TaobaoAilabAicloudTopDeviceDetailinfoGet(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceDetailinfoGetAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceDetailinfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
