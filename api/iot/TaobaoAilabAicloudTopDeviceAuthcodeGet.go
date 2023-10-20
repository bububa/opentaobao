package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceauthcodeget 获取设备授权码
// taobao.ailab.aicloud.top.device.authcode.get
//
// 获取设备授权码
func Taobaoailabaicloudtopdeviceauthcodeget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceauthcodegetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceauthcodegetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceauthcodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
