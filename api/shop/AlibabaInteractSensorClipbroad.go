package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// Alibabainteractsensorclipbroad Weex页面设置或读取剪切板
// alibaba.interact.sensor.clipbroad
//
// Weex页面设置或读取剪切板
func Alibabainteractsensorclipbroad(clt *core.SDKClient, req *shop.AlibabainteractsensorclipbroadAPIRequest, session string) (*shop.AlibabainteractsensorclipbroadAPIResponse, error) {
	var resp shop.AlibabainteractsensorclipbroadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
