package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdeviceopenidunbind openTaoBaoId解绑设备
// taobao.ailab.aicloud.top.device.openid.unbind
//
// openTaoBaoId解绑设备
func Taobaoailabaicloudtopdeviceopenidunbind(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdeviceopenidunbindAPIRequest, session string) (*iot.TaobaoailabaicloudtopdeviceopenidunbindAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdeviceopenidunbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
