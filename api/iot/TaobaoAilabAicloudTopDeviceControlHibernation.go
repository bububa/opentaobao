package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolhibernation 定时休眠
// taobao.ailab.aicloud.top.device.control.hibernation
//
// 定时休眠
func Taobaoailabaicloudtopdevicecontrolhibernation(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolhibernationAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolhibernationAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolhibernationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
