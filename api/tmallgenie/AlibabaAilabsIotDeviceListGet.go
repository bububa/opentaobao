package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAilabsIotDeviceListGet 获取iot设备列表
// alibaba.ailabs.iot.device.list.get
//
// 通过此接口获取用户名下的iot设备列表
func AlibabaAilabsIotDeviceListGet(clt *core.SDKClient, req *tmallgenie.AlibabaAilabsIotDeviceListGetAPIRequest, session string) (*tmallgenie.AlibabaAilabsIotDeviceListGetAPIResponse, error) {
	var resp tmallgenie.AlibabaAilabsIotDeviceListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
