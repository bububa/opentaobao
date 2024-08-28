package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttIotStatusPush iot设备状态变化通知接口
// youku.ott.iot.status.push
//
// ott iot设备状态通知
func YoukuOttIotStatusPush(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttIotStatusPushAPIRequest, resp *ottpay.YoukuOttIotStatusPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
