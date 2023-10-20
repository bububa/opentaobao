package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// AlibabaInteractSensorClipbroad Weex页面设置或读取剪切板
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
func AlibabaInteractSensorClipbroad(clt *core.SDKClient, req *shop.AlibabaInteractSensorClipbroadAPIRequest, session string) (*shop.AlibabaInteractSensorClipbroadAPIResponse, error) {
	var resp shop.AlibabaInteractSensorClipbroadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
