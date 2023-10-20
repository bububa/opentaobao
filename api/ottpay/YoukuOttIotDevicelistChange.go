package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttIotDevicelistChange iot设备列表变化接口
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
func YoukuOttIotDevicelistChange(clt *core.SDKClient, req *ottpay.YoukuOttIotDevicelistChangeAPIRequest, resp *ottpay.YoukuOttIotDevicelistChangeAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
