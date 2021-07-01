package aligenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aligenie"
)

/* AlibabaAilabsAligenieDeviceUnbind
设备解绑操作接口
alibaba.ailabs.aligenie.device.unbind

让开发者能根据设备ID进行解绑操作的接口 */
func AlibabaAilabsAligenieDeviceUnbind(clt *core.SDKClient, req *aligenie.AlibabaAilabsAligenieDeviceUnbindAPIRequest, session string) (*aligenie.AlibabaAilabsAligenieDeviceUnbindAPIResponse, error) {
	var resp aligenie.AlibabaAilabsAligenieDeviceUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
