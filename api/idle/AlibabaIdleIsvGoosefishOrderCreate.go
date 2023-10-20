package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleIsvGoosefishOrderCreate 闲鱼三方安康容器订单创建
// alibaba.idle.isv.goosefish.order.create
//
// 闲鱼三方安康容器订单创建
func AlibabaIdleIsvGoosefishOrderCreate(clt *core.SDKClient, req *idle.AlibabaIdleIsvGoosefishOrderCreateAPIRequest, session string) (*idle.AlibabaIdleIsvGoosefishOrderCreateAPIResponse, error) {
	var resp idle.AlibabaIdleIsvGoosefishOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
