package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicecontrolplayurl 点播url
// taobao.ailab.aicloud.top.device.control.playurl
//
// 点播url
func Taobaoailabaicloudtopdevicecontrolplayurl(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicecontrolplayurlAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicecontrolplayurlAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicecontrolplayurlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
