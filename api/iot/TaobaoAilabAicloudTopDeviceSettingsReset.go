package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Taobaoailabaicloudtopdevicesettingsreset 重置设备个性化设置
// taobao.ailab.aicloud.top.device.settings.reset
//
// 重置设备个性化设置
func Taobaoailabaicloudtopdevicesettingsreset(clt *core.SDKClient, req *iot.TaobaoailabaicloudtopdevicesettingsresetAPIRequest, session string) (*iot.TaobaoailabaicloudtopdevicesettingsresetAPIResponse, error) {
	var resp iot.TaobaoailabaicloudtopdevicesettingsresetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
