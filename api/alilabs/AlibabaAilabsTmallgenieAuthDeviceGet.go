package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabsTmallgenieAuthDeviceGet
获取设备详情
alibaba.ailabs.tmallgenie.auth.device.get

通过此接口获取设备详情 */
func AlibabaAilabsTmallgenieAuthDeviceGet(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceGetAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthDeviceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
