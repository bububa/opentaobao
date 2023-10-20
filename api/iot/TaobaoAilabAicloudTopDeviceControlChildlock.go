package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolchildlock 设备儿童锁
// taobao.ailab.aicloud.top.device.control.childlock
//
// 设备儿童锁
func Taobaoailabaicloudtopdevicecontrolchildlock(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolchildlockAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolchildlockAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolchildlockAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
