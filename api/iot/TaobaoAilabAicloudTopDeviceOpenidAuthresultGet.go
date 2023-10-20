package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceopenidauthresultget 获取openId设备授权码验证结果
// taobao.ailab.aicloud.top.device.openid.authresult.get
//
// 获取openId设备授权码验证结果
func Taobaoailabaicloudtopdeviceopenidauthresultget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceopenidauthresultgetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceopenidauthresultgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
