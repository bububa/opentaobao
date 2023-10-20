package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrollamp 台灯控制
// taobao.ailab.aicloud.top.device.control.lamp
//
// 台灯控制
func Taobaoailabaicloudtopdevicecontrollamp(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrollampAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrollampAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrollampAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
