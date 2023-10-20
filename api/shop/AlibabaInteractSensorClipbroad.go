package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaInteractSensorClipbroad Weex页面设置或读取剪切板
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
func AlibabaInteractSensorClipbroad(clt *core.SDKClient, req *shop.AlibabaInteractSensorClipbroadAPIRequest, resp *shop.AlibabaInteractSensorClipbroadAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
