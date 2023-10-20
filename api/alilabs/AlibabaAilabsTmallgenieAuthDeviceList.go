package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsTmallgenieAuthDeviceList 获取用户设备列表
// alibaba.ailabs.tmallgenie.auth.device.list
//
// 通过此接口获取用户绑定的设备信息列表
func AlibabaAilabsTmallgenieAuthDeviceList(clt *core.SDKClient, req *alilabs.AlibabaAilabsTmallgenieAuthDeviceListAPIRequest, session string) (*alilabs.AlibabaAilabsTmallgenieAuthDeviceListAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTmallgenieAuthDeviceListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
