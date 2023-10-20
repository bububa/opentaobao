package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicegetstatus 获取设备状态
// taobao.ailab.aicloud.top.device.getstatus
//
// 获取设备状态
func Taobaoailabaicloudtopdevicegetstatus(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicegetstatusAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicegetstatusAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicegetstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
