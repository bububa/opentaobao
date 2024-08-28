package alink

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// AlibabaAlinkDeviceUnbind 解绑设备
// alibaba.alink.device.unbind
//
// 阿里智能解绑设备
func AlibabaAlinkDeviceUnbind(ctx context.Context, clt *core.SDKClient, req *alink.AlibabaAlinkDeviceUnbindAPIRequest, resp *alink.AlibabaAlinkDeviceUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
