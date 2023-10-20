package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceauthresultget 获取设备授权码验证结果
// taobao.ailab.aicloud.top.device.authresult.get
//
// 获取设备授权码验证结果
func Taobaoailabaicloudtopdeviceauthresultget(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceauthresultgetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceauthresultgetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceauthresultgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
