package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

/* TaobaoAilabAicloudTopDeviceUnbind
解绑设备
taobao.ailab.aicloud.top.device.unbind

解绑设备 */
func TaobaoAilabAicloudTopDeviceUnbind(clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceUnbindAPIRequest, session string) (*iot.TaobaoAilabAicloudTopDeviceUnbindAPIResponse, error) {
	var resp iot.TaobaoAilabAicloudTopDeviceUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
