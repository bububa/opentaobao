package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceunbind 解绑设备
// taobao.ailab.aicloud.top.device.unbind
//
// 解绑设备
func Taobaoailabaicloudtopdeviceunbind(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceunbindAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceunbindAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
