package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkDeviceBind 绑定设备
// alibaba.alink.device.bind
//
// 阿里智能解绑设备
func AlibabaAlinkDeviceBind(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceBindAPIRequest, resp *alink.AlibabaAlinkDeviceBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
