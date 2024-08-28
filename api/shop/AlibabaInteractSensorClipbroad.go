package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaInteractSensorClipbroad Weex页面设置或读取剪切板
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
func AlibabaInteractSensorClipbroad(ctx context.Context, clt *core.SDKClient, req *shop.AlibabaInteractSensorClipbroadAPIRequest, resp *shop.AlibabaInteractSensorClipbroadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
