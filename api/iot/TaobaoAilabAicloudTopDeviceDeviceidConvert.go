package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicedeviceidconvert 开放设备id转换内部设备id
// taobao.ailab.aicloud.top.device.deviceid.convert
//
// 将开放设备id转换为内部设备id
func Taobaoailabaicloudtopdevicedeviceidconvert(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicedeviceidconvertAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicedeviceidconvertAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicedeviceidconvertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
