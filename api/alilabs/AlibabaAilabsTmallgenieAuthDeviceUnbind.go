package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabsTmallgenieAuthDeviceUnbind
解绑设备
alibaba.ailabs.tmallgenie.auth.device.unbind

通过此接口解绑天猫精灵设备 */
func AlibabaAilabsTmallgenieAuthDeviceUnbind(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceUnbindAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthDeviceUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
