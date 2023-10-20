package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvItemQuery 服务商闲鱼商品查询
// alibaba.idle.isv.item.query
//
// 服务商ISV闲鱼商品查询
func AlibabaIdleIsvItemQuery(clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemQueryAPIRequest, resp *idleisv.AlibabaIdleIsvItemQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
