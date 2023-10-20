package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudsmarthometopgenielinkreportdevice 零配方案上报设备
// taobao.ailab.aicloud.smarthome.top.genielink.reportdevice
//
// 零配方案中设备联网成功之后上报设备
func Taobaoailabaicloudsmarthometopgenielinkreportdevice(clt *core.SDKClient, req *iot.TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIRequest, session string) (*iot.TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIResponse, error) {
	var resp iot.TaobaoailabaicloudsmarthometopgenielinkreportdeviceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
