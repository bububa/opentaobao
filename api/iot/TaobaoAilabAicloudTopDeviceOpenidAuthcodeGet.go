package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceopenidauthcodeget 获取openid设备通用授权码
// taobao.ailab.aicloud.top.device.openid.authcode.get
//
// 获取openid设备通用授权码
func Taobaoailabaicloudtopdeviceopenidauthcodeget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceopenidauthcodegetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceopenidauthcodegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
