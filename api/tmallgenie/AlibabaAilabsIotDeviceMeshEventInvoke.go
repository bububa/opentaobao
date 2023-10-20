package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaailabsiotdevicemesheventinvoke 弹内设备中心事件调用
// alibaba.ailabs.iot.device.mesh.event.invoke
//
// 弹内设备中心事件调用
func Alibabaailabsiotdevicemesheventinvoke(clt *core.SDKClient, req *tmallgenie.AlibabaailabsiotdevicemesheventinvokeAPIRequest, session string) (*tmallgenie.AlibabaailabsiotdevicemesheventinvokeAPIResponse, error) {
	var resp tmallgenie.AlibabaailabsiotdevicemesheventinvokeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
