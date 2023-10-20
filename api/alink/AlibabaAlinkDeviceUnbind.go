package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkDeviceUnbind 解绑设备
// alibaba.alink.device.unbind
//
// 阿里智能解绑设备
func AlibabaAlinkDeviceUnbind(clt *core.SDKClient, req *alink.AlibabaAlinkDeviceUnbindAPIRequest, resp *alink.AlibabaAlinkDeviceUnbindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
