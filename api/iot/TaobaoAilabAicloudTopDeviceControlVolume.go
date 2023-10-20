package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolvolume 设备音量
// taobao.ailab.aicloud.top.device.control.volume
//
// 设备音量
func Taobaoailabaicloudtopdevicecontrolvolume(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolvolumeAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolvolumeAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolvolumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
