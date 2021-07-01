package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

/* AlibabaAilabsTvsDeviceList
获取TVS设备列表
alibaba.ailabs.tvs.device.list

获取用户所绑定的TVS设备列表 */
func AlibabaAilabsTvsDeviceList(clt *core.SDKClient, req *alilabs.AlibabaAilabsTvsDeviceListAPIRequest, session string) (*alilabs.AlibabaAilabsTvsDeviceListAPIResponse, error) {
	var resp alilabs.AlibabaAilabsTvsDeviceListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
