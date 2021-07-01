package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGet
根据三方ID查询设备注册激活信息
alibaba.ailabs.tmallgenie.auth.device.withdeviceid.get

根据三方ID查询设备注册激活信息 */
func AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGet(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthDeviceWithdeviceidGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
