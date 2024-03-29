package alilabs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alilabs"
)

// AlibabaAilabsIotDeviceListUpdateNotify 设备列表更新通知
// alibaba.ailabs.iot.device.list.update.notify
//
// 用于人工智能实验室IoT合作厂商上报三方接入IoT设备列表更新时的通知
func AlibabaAilabsIotDeviceListUpdateNotify(clt *core.SDKClient, req *alilabs.AlibabaAilabsIotDeviceListUpdateNotifyAPIRequest, resp *alilabs.AlibabaAilabsIotDeviceListUpdateNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
