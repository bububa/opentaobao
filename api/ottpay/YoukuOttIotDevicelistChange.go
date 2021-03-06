package ottpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttIotDevicelistChange iot设备列表变化接口
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
func YoukuOttIotDevicelistChange(clt *core.SDKClient, req *ottpay.YoukuOttIotDevicelistChangeAPIRequest, session string) (*ottpay.YoukuOttIotDevicelistChangeAPIResponse, error) {
	var resp ottpay.YoukuOttIotDevicelistChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
