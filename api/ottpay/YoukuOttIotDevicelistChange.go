package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttIotDevicelistChange iot设备列表变化接口
// youku.ott.iot.devicelist.change
//
// iot设备列表变化接口
func YoukuOttIotDevicelistChange(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttIotDevicelistChangeAPIRequest, resp *ottpay.YoukuOttIotDevicelistChangeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
