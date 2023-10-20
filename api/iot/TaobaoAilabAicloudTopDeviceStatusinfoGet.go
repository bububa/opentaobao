package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicestatusinfoget 获取设备状态信息
// taobao.ailab.aicloud.top.device.statusinfo.get
//
// 获取设备状态信息
func Taobaoailabaicloudtopdevicestatusinfoget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicestatusinfogetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicestatusinfogetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicestatusinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
