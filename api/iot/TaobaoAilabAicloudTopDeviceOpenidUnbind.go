package iot

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// TaobaoAilabAicloudTopDeviceOpenidUnbind openTaoBaoId解绑设备
// taobao.ailab.aicloud.top.device.openid.unbind
//
// openTaoBaoId解绑设备
func TaobaoAilabAicloudTopDeviceOpenidUnbind(ctx context.Context, clt *core.SDKClient, req *iot.TaobaoAilabAicloudTopDeviceOpenidUnbindAPIRequest, resp *iot.TaobaoAilabAicloudTopDeviceOpenidUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
