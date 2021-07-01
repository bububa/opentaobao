package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

/* AlibabaInteractWindvaneCall
Weex页面调用windvane
alibaba.interact.windvane.call

客户端鉴权使用，实际不会发送或接收数据 */
func AlibabaInteractWindvaneCall(clt *core.SDKClient, req *shop.AlibabaInteractWindvaneCallAPIRequest, session string) (*shop.AlibabaInteractWindvaneCallAPIResponse, error) {
	var resp shop.AlibabaInteractWindvaneCallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
